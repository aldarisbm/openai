package main

import (
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

// setup is a function that sets up the environment, and returns a ChatBot.
func setup() *ChatBot {
	// Load the .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the token from the .env file.
	token := os.Getenv("TOKEN")

	// Get the system context from the .env file.
	systemContext := os.Getenv("SYSTEM_CONTEXT")

	// If the token is empty, get it from the user.
	if token == "" {
		token = getInput("Provide your OpenAI token")
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
			MaxMessages: 2,
		},
		client: c,
	}
}
