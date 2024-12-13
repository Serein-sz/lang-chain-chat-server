package model

import "github.com/tmc/langchaingo/llms"

type Chat struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    llms.ChatMessageType `json:"role"`
	Content string               `json:"content"`
}

type MessageVo struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}
