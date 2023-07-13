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

func getModelFromEnv() string {
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

func getUserInput(s string) string {
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
