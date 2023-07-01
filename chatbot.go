package main

import (
	"github.com/sashabaranov/go-openai"
)

const (
	user      = "user"
	assistant = "assistant"
)

type ChatBot struct {
	apiToken    string
	chatContext *ChatContext
	client      *openai.Client
	model       string
}

func (b ChatBot) getRequest(prompt string) openai.ChatCompletionRequest {
	b.saveMessage(user, prompt)
	req := b.getMessage()
	b.appendMessages(&req)
	return req
}

func (b ChatBot) saveMessage(role, msg string) {
	b.chatContext.Messages = append(b.chatContext.Messages, openai.ChatCompletionMessage{
		Role:    role,
		Content: msg,
	})
}

func (b ChatBot) ClearMessages() {
	b.chatContext.Messages = []openai.ChatCompletionMessage{}
}

func (b ChatBot) getMessage() openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: b.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "You're a helpful assistant, you are very concise.",
			},
		},
		N:           1,
		MaxTokens:   0,
		Temperature: 0.2,
		Stream:      true,
	}
}

func (b ChatBot) appendMessages(req *openai.ChatCompletionRequest) {
	if len(b.chatContext.Messages) <= b.chatContext.MaxPriorMessages {
		req.Messages = append(req.Messages, b.chatContext.Messages...)
		return
	} else {
		req.Messages = append(req.Messages, b.chatContext.Messages[:b.chatContext.MaxPriorMessages]...)
		b.chatContext.Messages = b.chatContext.Messages[b.chatContext.MaxPriorMessages:]
	}
}
