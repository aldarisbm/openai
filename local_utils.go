package main

import (
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strconv"
)

// setup is a function that sets up the environment, and returns a ChatBot.
func setup() *ChatBot {
	// Load the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	systemContext := os.Getenv("SYSTEM_CONTEXT")
	maxPriorMessagesString := os.Getenv("MAX_PRIOR_MESSAGES")
	maxPriorMessages, err := strconv.Atoi(maxPriorMessagesString)
	if err != nil {
		panic(err)
	}
	chatContextRole := os.Getenv("CHAT_CONTEXT_ROLE")

	// If the token is empty, get it from the user.
	if token == "" {
		token = getInput("Provide your OpenAI token")
	}
	if err = validateToken(token); err != nil {
		panic(err)
	}

	// If the system context is empty, get it from the user.
	if systemContext == "" {
		systemContext = getInput("System context")
	}

	// Create a new ChatBot.
	c := openai.NewClient(token)

	// Return the ChatBot.
	return &ChatBot{
		apiToken:      token,
		systemContext: systemContext,
		chatContext: ChatContext{
			Role:             chatContextRole,
			MaxPriorMessages: maxPriorMessages,
		},
		client: c,
	}
}
