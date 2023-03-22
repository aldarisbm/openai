package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strconv"
	"strings"
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
		chatContext: &ChatContext{
			Role:             chatContextRole,
			MaxPriorMessages: maxPriorMessages,
		},
		client: c,
	}
}

func getSummaryBetweenThreeBrackets(message string) string {
	sum := ""
	start := strings.Index(message, "[[[")
	end := strings.Index(message, "]]]")
	if start != -1 && end != -1 {
		sum = message[start+3 : end]
	}
	return fmt.Sprintf("[[[%s]]]", sum)
}

func getInput(s string) string {
	fmt.Print(fmt.Sprintf("%s:", s))
	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	trimmedInput := strings.TrimSpace(token)

	return trimmedInput
}

func validateToken(token string) error {
	if len(token) != 51 {
		return InvalidTokenError
	}
	return nil
}
