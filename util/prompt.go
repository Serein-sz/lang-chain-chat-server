package util

import "github.com/tmc/langchaingo/prompts"

func CreateSystemPrompt() prompts.ChatPromptTemplate {
	return prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		prompts.NewSystemMessagePromptTemplate("{{.system}}", []string{"system"}),
	})
}
