package main

import (
	"os"
	"strconv"
)

type Environment struct {
	Token            string
	MaxPriorMessages int
	Model            string
	SystemPrompt     string
}

func GetEnvironment() *Environment {
	token := os.Getenv("TOKEN")
	maxPriorMessagesString := os.Getenv("MAX_PRIOR_MESSAGES")
	maxPriorMessages, err := strconv.Atoi(maxPriorMessagesString)
	if err != nil {
		maxPriorMessages = MaxPriorMessages
	}
	systemPrompt := os.Getenv("SYSTEM_PROMPT")
	if systemPrompt == "" {
		systemPrompt = DefaultPrompt
	}

	if token == "" {
		token = getUserInput("Provide your OpenAI token")
	}
	if err = validateToken(token); err != nil {
		panic(err)
	}

	model := getModelFromEnv()
	return &Environment{
		Token:            token,
		MaxPriorMessages: maxPriorMessages,
		Model:            model,
		SystemPrompt:     systemPrompt,
	}
}
