package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"io"
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

	message := make(chan string, 20)

	go func() {
		_, err := llm.GenerateContent(c, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			message <- string(chunk)
			return nil
		}))
		defer close(message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}()

	c.Stream(func(w io.Writer) bool {
		item, ok := <-message
		c.SSEvent("chat", item)
		return ok
	})

}
