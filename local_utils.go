package main

import (
	"bufio"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strconv"
	"strings"
)

func getModelFromEnv() string {
	enableGPT4String := os.Getenv("ENABLE_GPT_4")
	isChatGPT4Enabled, err := strconv.ParseBool(enableGPT4String)
	if err != nil {
		panic(err)
	}
	if isChatGPT4Enabled {
		return openai.GPT4o
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
