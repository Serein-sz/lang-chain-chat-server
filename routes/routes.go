package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(r *gin.Engine) {
	log.Println("Registering all routes...")
	ChatRoutes(r)
}
