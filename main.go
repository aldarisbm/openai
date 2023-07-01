package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {
	loadEnvFile()
	bot := New()
	ctx := context.Background()

	fmt.Printf("model: %s\nmax message history: %d\n\n", bot.model, bot.chatContext.MaxPriorMessages)

	for {
		prompt := getUserInput("Prompt")
		if prompt == "" {
			continue
		}
		if prompt == "clear" || prompt == "c" {
			bot.ClearMessages()
			fmt.Println("Messages cleared...")
			continue
		}

		req := bot.GetRequest(prompt)

		stream, err := bot.client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
		}

		if stream.GetResponse().StatusCode != http.StatusOK {
			body, err := io.ReadAll(stream.GetResponse().Body)
			if err != nil {
				fmt.Printf("error reading body: %v\n", err)
				continue
			}
			fmt.Printf("non 200 resp: %v\n", string(body))
		}
		defer stream.Close()

		var lastMessage string
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) && lastMessage != "" {
				bot.saveMessageToContext(assistant, lastMessage)
				lastMessage = ""
				fmt.Println()
				break
			}
			if err != nil {
				fmt.Printf("Stream error: %v\n", err)
			}
			incomingStr := response.Choices[0].Delta.Content
			lastMessage += incomingStr
			fmt.Print(incomingStr)
		}
	}
}
