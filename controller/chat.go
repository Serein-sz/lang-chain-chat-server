package controller

import (
	"lang-chain-chat-server/model"
	"lang-chain-chat-server/service"
	"lang-chain-chat-server/util"
	"net/http"

	"github.com/tmc/langchaingo/llms"

	"github.com/gin-gonic/gin"
)

type Chat struct{}

func (ch *Chat) DoChat(c *gin.Context) {

	var body model.Chat

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	history := service.GenerateHistory(body.Messages)

	messages := make([]llms.MessageContent, 0)

	prompt, _ := util.CreateSystemPrompt().Format(map[string]interface{}{
		"system": `
			You are a chinese programming expert and you will answer my questions in markdown format.
			`,
	})

	messages = append(messages, llms.TextParts(llms.ChatMessageTypeSystem, prompt))

	for _, chatMessage := range history {
		messages = append(messages, llms.TextParts(chatMessage.GetType(), chatMessage.GetContent()))
		chatMessage.GetType()
	}
	llm, err := util.CreateLlm(body.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	service.HandleSse(c, llm, messages)

}
