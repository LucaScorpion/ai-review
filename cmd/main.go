package main

import (
	"ai-reviewer/internal/openai"
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("No OPENAI_API_KEY found in environment.")
		os.Exit(1)
	}

	client := openai.NewClient(apiKey, openai.Gpt4TurboPreview)

	res := client.CreateCompletion([]openai.Message{
		{
			Role:    openai.RoleSystem,
			Content: "You are a helpful AI assistant named EZ-review.",
		},
		{
			Role:    openai.RoleUser,
			Content: "Hi there! What is your name?",
		},
	})

	fmt.Println(res.Content)
}
