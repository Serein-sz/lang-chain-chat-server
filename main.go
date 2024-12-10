package main

import (
	"lang-chain-chat-server/middleware"
	"lang-chain-chat-server/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	routes.RegisterAllRoutes(r)

	err := r.Run(":8000")
	if err != nil {
		log.Fatal("server start error: {}", err.Error())
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server is shutting down...")
}
