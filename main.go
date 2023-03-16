package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	token := getInput("Provide your OpenAI token")
	err := validateToken(token)
	if err != nil {
		log.Fatalf("validating token: %s\n", err)
	}

	systemContext := getInput("System context")
	c := openai.NewClient(token)
	ctx := context.Background()

	for {
		prompt := getInput("Prompt")
		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{{
				Role:    "user",
				Content: prompt,
			}},
			N:           1,
			MaxTokens:   1024,
			Temperature: 0.6,
			Stream:      true,
		}
		if systemContext != "" {
			req.Messages = append(req.Messages, openai.ChatCompletionMessage{
				Role:    "system",
				Content: systemContext,
			})
		} else {
			req.Messages = append(req.Messages, openai.ChatCompletionMessage{
				Role:    "system",
				Content: "You're GoGPT, a great golang code assistant.",
			})
		}

		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
		}

		if stream.GetResponse().StatusCode == http.StatusBadRequest {
			fmt.Printf("check: %v\n", stream.GetResponse())
		}

		defer stream.Close()

		LineCharLimit := 100
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
			fmt.Printf(incomingStr)
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
