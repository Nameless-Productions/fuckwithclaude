package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env")
	}
	claudeApiKey := os.Getenv("CLAUDE")
	if claudeApiKey == "" {
		log.Fatal("No CLAUDE env variable exists")
	}


	fmt.Println("Fuck with Claude")

	client := anthropic.NewClient(
		option.WithAPIKey(claudeApiKey),
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input: ")
	scanner.Scan()
	msg := scanner.Text()

	fmt.Print("Times: ")
	scanner.Scan()
	times := scanner.Text()

	timesInt, err := strconv.Atoi(times)
	if err != nil {
		log.Fatalln("NaN")
	}

	msgToClaude := msg

	for range timesInt {
		res, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
			Model: anthropic.ModelClaudeHaiku4_5,
			MaxTokens: 1024,
			Messages: []anthropic.MessageParam{
				anthropic.NewUserMessage(anthropic.NewTextBlock(msgToClaude)),
			},
		})

		if err != nil {
			log.Fatal(err)
		}
		
		msgToClaude = res.Content[0].Text
	}

	fmt.Println(msgToClaude)
}