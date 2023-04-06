package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const CharLimit = 150

func main() {
	bot := setup()
	ctx := context.Background()

	fmt.Println("Welcome to the OpenAI Chat Completion")
	fmt.Printf("You're using: %s\n", bot.model)
	for {
		prompt := getInput("Prompt")
		if prompt == "" {
			continue
		}
		if prompt == "clear" || prompt == "c" {
			bot.ClearMessages()
			fmt.Println("Messages cleared...")
			continue
		}
		req := bot.getRequest(prompt)

		stream, err := bot.client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
		}

		if stream.GetResponse().StatusCode != http.StatusOK {
			body, _ := io.ReadAll(stream.GetResponse().Body)
			fmt.Printf("non 200 resp: %v\n", string(body))
		}

		defer stream.Close()

		LineCharLimit := CharLimit
		CurrentLineLength := 0
		var wordsInLine []string

		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				bot.processEntireMessage()
				bot.lastEntireMessage = ""
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
