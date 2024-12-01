package routes

import "github.com/gin-gonic/gin"

func RegisterAllRoutes(r *gin.Engine) {
	ChatRoutes(r)
}
