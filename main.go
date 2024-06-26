package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	bot := New()
	ctx := context.Background()

	fmt.Printf("model: %s\nsystem prompt: %s\nmax message history: %d\n\n", bot.model, bot.GetPrompt(), bot.chatContext.MaxPriorMessages)

	for {
		prompt := getUserInput("Prompt")
		switch prompt {
		case "history":
			bot.PrintHistory()
		case "":
			continue
		case "q", "quit", "exit":
			fmt.Println("Exiting...")
			return
		case "clear", "c":
			bot.ClearMessages()
			fmt.Println("Messages cleared...")
			continue
		default:
			process(ctx, bot, prompt)
		}
	}
}

func process(ctx context.Context, bot *ChatBot, prompt string) {
	req := bot.GetRequest(prompt)

	stream, err := bot.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
	}
	defer stream.Close()

	var lastMessage string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) && lastMessage != "" {
			bot.saveMessageToContext(Assistant, lastMessage)
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
