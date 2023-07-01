package main

import "github.com/sashabaranov/go-openai"

// ChatContext represents a chat context.
type ChatContext struct {
	Messages         []openai.ChatCompletionMessage `json:"messages"`
	MaxPriorMessages int                            `json:"max_prior_messages"`
}
