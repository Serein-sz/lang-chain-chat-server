package routes

import (
	"github.com/gin-gonic/gin"
	"lang-chain-chat-server/controller"
)

func ChatRoutes(r *gin.Engine) {
	chatRouters := r.Group("/chat")
	{
		chat := controller.Chat{}
		chatRouters.POST("/", chat.DoChat)
	}
}
