package main

import (
	"fmt"
	"lang-chain-chat-server/middleware"
	"lang-chain-chat-server/routes"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())
	
	routes.RegisterAllRoutes(r)

	r.Run(":8000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
}
