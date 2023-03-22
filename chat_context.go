package main

import "github.com/sashabaranov/go-openai"

// ChatContext represents a chat context.
type ChatContext struct {
	// Messages is a list of messages that have been exchanged so far.
	Messages []openai.ChatCompletionMessage `json:"messages"`

	//Role is the role of the chat context. it defaults to "user"
	Role string

	// MaxPriorMessages is the maximum number of prior messages to include in the chat context.
	MaxPriorMessages int `json:"max_prior_messages"`
}
