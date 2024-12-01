package util

import (
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/ollama"
	"net/http"
)

func CreateOllama(c *gin.Context, modelName string) *ollama.LLM {
	llm, err := ollama.New(ollama.WithModel(modelName))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return nil
	}
	return llm
}
