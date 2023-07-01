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
	MaxPriorMessages int
	Model            string
}

func New() *ChatBot {
	env := NewEnvironment()
	return &ChatBot{
		apiToken: env.Token,
		chatContext: &ChatContext{
			MaxPriorMessages: env.MaxPriorMessages,
		},
		client: openai.NewClient(env.Token),
		model:  env.Model,
	}
}

func loadEnvFile() {
	envFileLoc := os.Getenv("OPENAPI_ENV_FILE")
	if envFileLoc == "" {
		panic("ENV_FILE is not set")
	}

	err := godotenv.Load(envFileLoc)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewEnvironment() *Environment {
	token := os.Getenv("TOKEN")
	maxPriorMessagesString := os.Getenv("MAX_PRIOR_MESSAGES")
	maxPriorMessages, err := strconv.Atoi(maxPriorMessagesString)
	if err != nil {
		maxPriorMessages = MaxPriorMessages
	}

	if token == "" {
		token = getInput("Provide your OpenAI token")
	}
	if err = validateToken(token); err != nil {
		panic(err)
	}

	model := getModel()
	return &Environment{
		Token:            token,
		MaxPriorMessages: maxPriorMessages,
		Model:            model,
	}
}

func getModel() string {
	enableGPT4String := os.Getenv("ENABLE_GPT_4")
	isChatGPT4Enabled, err := strconv.ParseBool(enableGPT4String)
	if err != nil {
		panic(err)
	}
	if isChatGPT4Enabled {
		return openai.GPT4
	}
	return openai.GPT3Dot5Turbo

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
