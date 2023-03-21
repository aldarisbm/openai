package main

// ChatContext represents a chat context.
type ChatContext struct {
	// Messages is a list of messages that have been exchanged so far.
	Messages []ChatCompletionMessage `json:"messages"`

	// MaxMessages
	MaxMessages int `json:"max_messages"`
}

// ChatCompletionMessage represents prior messages in a chat context.
type ChatCompletionMessage struct {
	// Role is the role of the message sender. it defaults to "system"
	Role string `json:"role"`

	Message string `json:"message"`
}
