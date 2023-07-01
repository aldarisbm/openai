package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

const (
	user      = "user"
	assistant = "assistant"
)

// ChatBot represents a chatbot.
type ChatBot struct {
	apiToken     string
	chatContext  *ChatContext
	client       *openai.Client
	model        string
	systemPrompt string
}

// ChatContext represents a chat context.
type ChatContext struct {
	Messages         []openai.ChatCompletionMessage
	MaxPriorMessages int
}

func New() *ChatBot {
	env := GetEnvironment()
	return &ChatBot{
		apiToken: env.Token,
		chatContext: &ChatContext{
			MaxPriorMessages: env.MaxPriorMessages,
		},
		systemPrompt: env.SystemPrompt,
		client:       openai.NewClient(env.Token),
		model:        env.Model,
	}
}

func (b ChatBot) GetPrompt() string {
	return b.systemPrompt
}

func (b ChatBot) GetRequest(prompt string) openai.ChatCompletionRequest {
	b.saveMessageToContext(user, prompt)
	req := b.getBaseMessage()
	b.appendMessages(&req)
	return req
}
func (b ChatBot) ClearMessages() {
	b.chatContext.Messages = []openai.ChatCompletionMessage{}
}

func (b ChatBot) saveMessageToContext(role, msg string) {
	b.chatContext.Messages = append(b.chatContext.Messages, openai.ChatCompletionMessage{
		Role:    role,
		Content: msg,
	})
	if len(b.chatContext.Messages) > b.chatContext.MaxPriorMessages {
		b.chatContext.Messages = b.chatContext.Messages[len(b.chatContext.Messages)-b.chatContext.MaxPriorMessages:]
	}
}

func (b ChatBot) getBaseMessage() openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: b.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: b.systemPrompt,
			},
		},
		N:           1,
		MaxTokens:   0,
		Temperature: 0.2,
		Stream:      true,
	}
}

func (b ChatBot) appendMessages(req *openai.ChatCompletionRequest) {
	req.Messages = append(req.Messages, b.chatContext.Messages...)
}

func (b ChatBot) PrintHistory() {
	fmt.Println("*************************")
	fmt.Printf("system prompt: %s\n", b.GetPrompt())
	for _, msg := range b.chatContext.Messages {
		fmt.Printf("%s: %s\n", msg.Role, msg.Content)
	}
	fmt.Println("*************************")
}
