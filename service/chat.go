package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"io"
	"lang-chain-chat-server/model"
	"net/http"
)

func GenerateHistory(messages []model.Message) []llms.ChatMessage {
	history := memory.NewChatMessageHistory()
	for _, message := range messages {
		if message.Role == llms.ChatMessageTypeAI {
			_ = history.AddAIMessage(context.Background(), message.Content)
		}
		if message.Role == llms.ChatMessageTypeHuman {
			_ = history.AddUserMessage(context.Background(), message.Content)
		}
	}
	chatMessages, _ := history.Messages(context.Background())
	return chatMessages
}

func HandleSse(c *gin.Context, llm *ollama.LLM, messageContents []llms.MessageContent) {

	message := make(chan string, 20)

	go func() {
		_, err := llm.GenerateContent(c, messageContents, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
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
		content, _ := json.Marshal(model.MessageVo{
			Key:     "1",
			Message: item,
		})
		c.SSEvent("chat", string(content))
		return ok
	})
}
