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

const CharLimit = 150

type ChatBot struct {
	apiToken          string
	systemContext     string
	chatContext       ChatContext
	client            *openai.Client
	lastEntireMessage string
}

func (b ChatBot) getRequest(prompt string) openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: b.systemContext,
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		N:           1,
		MaxTokens:   1024,
		Temperature: 0.5,
		Stream:      true,
	}

	if len(b.chatContext.Messages) > b.chatContext.MaxPriorMessages {
		// Remove the oldest message.
		b.chatContext.Messages = b.chatContext.Messages[1:]
	}

	for _, message := range b.chatContext.Messages {
		req.Messages = append(req.Messages, message)
	}
	return req
}

func (b ChatBot) processEntireMessage() {
	currentSummary := getSummaryBetweenBrackets(b.lastEntireMessage)
	b.chatContext.Messages = append(b.chatContext.Messages, openai.ChatCompletionMessage{
		Role:    b.chatContext.Role,
		Content: currentSummary,
	})
	b.lastEntireMessage = ""
}

func getSummaryBetweenBrackets(message string) string {
	sum := ""
	start := strings.Index(message, "[[[")
	end := strings.Index(message, "]]]")
	if start != -1 && end != -1 {
		sum = message[start+3 : end]
	}
	return sum
}

func main() {
	bot := setup()
	ctx := context.Background()

	fmt.Printf("system:%s\n", bot.systemContext)
	for {
		prompt := getInput("Prompt")
		if prompt == "" {
			continue
		}
		req := bot.getRequest(prompt)

		stream, err := bot.client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
		}

		if stream.GetResponse().StatusCode == http.StatusBadRequest {
			fmt.Printf("bad request: %v\n", stream.GetResponse())
		}

		if stream.GetResponse().StatusCode != http.StatusOK {
			fmt.Printf("non 200 resp: %v\n", stream.GetResponse())
		}
		defer stream.Close()

		LineCharLimit := CharLimit
		CurrentLineLength := 0
		var wordsInLine []string

		for {

			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				bot.processEntireMessage()
				fmt.Println()
				break
			}
			if err != nil {
				fmt.Printf("Stream error: %v\n", err)
			}
			incomingStr := response.Choices[0].Delta.Content
			bot.lastEntireMessage += incomingStr
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
