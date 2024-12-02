package util

import (
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/ollama"
	"net/http"
)

func CreateOllama(c *gin.Context, modelName string) *ollama.LLM {
	llm, err := ollama.New(ollama.WithModel(modelName), ollama.WithServerURL("http://112.125.89.224:11434"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return nil
	}
	return llm
}
