package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const CharLimit = 150

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TOKEN")
	systemContext := os.Getenv("SYSTEM_CONTEXT")

	if token == "" {
		token = getInput("Provide your OpenAI token")
	}
	err = validateToken(token)
	if err != nil {
		log.Fatalf("validating token: %s\n", err)
	}

	if systemContext == "" {
		systemContext = getInput("System context")
	}

	c := openai.NewClient(token)
	ctx := context.Background()

	fmt.Printf("system:%s\n", systemContext)
	for {
		prompt := getInput("Prompt")
		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: prompt,
				},
				{
					Role:    "system",
					Content: systemContext,
				},
			},
			N:           1,
			MaxTokens:   1024,
			Temperature: 0.6,
			Stream:      true,
		}

		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
		}

		if stream.GetResponse().StatusCode == http.StatusBadRequest {
			fmt.Printf("check: %v\n", stream.GetResponse())
		}

		if stream.GetResponse().StatusCode != http.StatusOK {
			fmt.Printf("check: %v\n", stream.GetResponse())
		}
		defer stream.Close()

		LineCharLimit := CharLimit
		CurrentLineLength := 0
		var wordsInLine []string

		for {

			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println()
				break
			}
			if err != nil {
				fmt.Printf("Stream error: %v\n", err)
			}
			incomingStr := response.Choices[0].Delta.Content
			CurrentLineLength += len(incomingStr)
			wordsInLine = append(wordsInLine, incomingStr)
			fmt.Print(incomingStr)
			if CurrentLineLength >= LineCharLimit {
				fmt.Println()
				CurrentLineLength = 0
			}
		}
	}
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
		return fmt.Errorf("invalid token length")
	}

	return nil
}
