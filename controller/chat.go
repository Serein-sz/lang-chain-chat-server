package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"lang-chain-chat-server/model"
	"lang-chain-chat-server/util"
	"net/http"
)

type Chat struct{}

func (ch *Chat) DoChat(c *gin.Context) {
	var body model.Chat

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	prompt := util.CreatePrompt()

	data := map[string]interface{}{
		"text": body.Text,
	}

	msg, _ := prompt.FormatMessages(data)

	content := []llms.MessageContent{
		llms.TextParts(msg[0].GetType(), msg[0].GetContent()),
		llms.TextParts(msg[1].GetType(), msg[1].GetContent()),
	}

	llm := util.CreateOllama(c, "qwen2.5-coder:0.5b")

	resp, err := llm.GenerateContent(c, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": resp.Choices[0].Content,
	})
}
