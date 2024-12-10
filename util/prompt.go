package util

import "github.com/tmc/langchaingo/prompts"

func CreatePrompt() prompts.ChatPromptTemplate {
	return prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		prompts.NewSystemMessagePromptTemplate("{{.system}}", []string{"system"}),
		prompts.NewHumanMessagePromptTemplate("{{.text}}", []string{"text"}),
	})
}
