package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type ChatBot struct {
	apiToken          string
	systemContext     string
	chatContext       *ChatContext
	client            *openai.Client
	lastEntireMessage string
	model             string
}

func (b ChatBot) getRequest(prompt string) openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model: b.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: b.systemContext,
			},
		},
		N:           1,
		MaxTokens:   0,
		Temperature: 0.5,
		Stream:      true,
	}

	if len(b.chatContext.Messages) > b.chatContext.MaxPriorMessages {
		// Remove the oldest message.
		b.chatContext.Messages = b.chatContext.Messages[1:]
	}

	for _, message := range b.chatContext.Messages {
		message.Content = fmt.Sprintf("%s", message.Content)
		req.Messages = append(req.Messages, message)
	}
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: prompt,
	})

	return req
}

func (b ChatBot) processEntireMessage() {
	currentSummary := getSummaryBetweenThreeBrackets(b.lastEntireMessage)
	b.chatContext.Messages = append(b.chatContext.Messages, openai.ChatCompletionMessage{
		Role:    b.chatContext.Role,
		Content: currentSummary,
	})
}
