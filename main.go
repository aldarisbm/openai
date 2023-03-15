package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please provide your token:")
	reader := bufio.NewReader(os.Stdin)
	token, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	trimmedToken := strings.TrimSpace(token)
	err = validateToken(trimmedToken)
	if err != nil {
		log.Fatalf("validating token: %s\n", err)
	}
	fmt.Printf("token length:%d\n", len(trimmedToken))
	c := openai.NewClient(token)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("CompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}

func validateToken(token string) error {
	if len(token) != 51 {
		return fmt.Errorf("invalid token length")
	}

	return nil
}
