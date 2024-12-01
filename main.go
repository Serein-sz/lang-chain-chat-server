package main

import (
	"github.com/gin-gonic/gin"
	"lang-chain-chat-server/routes"
	"net/http"
)

func main() {
	r := gin.New()
	routes.RegisterAllRoutes(r)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Hello World",
		})
	})

	r.Run(":8000")
}
