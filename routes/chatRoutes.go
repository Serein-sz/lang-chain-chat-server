package routes

import (
	"lang-chain-chat-server/controller"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(r *gin.Engine) {
	chatRouters := r.Group("/chat")
	chat := controller.Chat{}
	{
		chatRouters.POST("/", chat.DoChat)
	}
}
