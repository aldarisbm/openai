package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed .env
var envFile embed.FS

// setup is a function that sets up the environment, and returns a ChatBot.
func setup() *ChatBot {
	populateEnvironment()

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

func populateEnvironment() {
	// Read the embedded .env file
	data, err := fs.ReadFile(envFile, ".env")
	if err != nil {
		log.Fatalf("Error reading embedded .env file: %v", err)
	}

	// Use a bytes.Buffer to create an io.Reader for the .env data
	envReader := bytes.NewReader(data)
	// Load the .env file.
	envVars, err := godotenv.Parse(envReader)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for key, value := range envVars {
		os.Setenv(key, value)
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
