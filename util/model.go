package util

import (
	"github.com/tmc/langchaingo/llms/ollama"
)

func CreateLlm(modelName string) (*ollama.LLM, error) {
	return ollama.New(ollama.WithModel(modelName), ollama.WithServerURL("http://112.125.89.224:11434"))
}
