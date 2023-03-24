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

const MaxPriorMessages = 5

type Environment struct {
	Token            string
	SystemContext    string
	MaxPriorMessages int
	ChatContextRole  string
}

// setup is a function that sets up the environment, and returns a ChatBot.
func setup() *ChatBot {
	populateEnvironment()
	Env := NewEnvironment()
	return &ChatBot{
		apiToken:      Env.Token,
		systemContext: Env.SystemContext,
		chatContext: &ChatContext{
			Role:             Env.ChatContextRole,
			MaxPriorMessages: Env.MaxPriorMessages,
		},
		client: openai.NewClient(Env.Token),
	}
}

func populateEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewEnvironment() *Environment {
	token := os.Getenv("TOKEN")
	systemContext := os.Getenv("SYSTEM_CONTEXT")
	maxPriorMessagesString := os.Getenv("MAX_PRIOR_MESSAGES")
	maxPriorMessages, err := strconv.Atoi(maxPriorMessagesString)
	if err != nil {
		maxPriorMessages = MaxPriorMessages
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
	return &Environment{
		Token:            token,
		SystemContext:    systemContext,
		MaxPriorMessages: maxPriorMessages,
		ChatContextRole:  chatContextRole,
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
