package main

import (
	"os"
	"strconv"
)

const MaxPriorMessages = 5

type Environment struct {
	Token            string
	MaxPriorMessages int
	Model            string
}

func GetEnvironment() *Environment {
	token := os.Getenv("TOKEN")
	maxPriorMessagesString := os.Getenv("MAX_PRIOR_MESSAGES")
	maxPriorMessages, err := strconv.Atoi(maxPriorMessagesString)
	if err != nil {
		maxPriorMessages = MaxPriorMessages
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
	}
}
