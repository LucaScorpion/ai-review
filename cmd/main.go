package main

import (
	"ai-reviewer/internal/git"
	"ai-reviewer/internal/openai"
	"ai-reviewer/internal/prompt"
	"ai-reviewer/internal/util"
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	// Check if the OpenAI API key is set.
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		errAndExit("No OPENAI_API_KEY found in environment.")
	}

	// Check if we have a target branch.
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		errAndExit("No target diff branch given.")
	}

	// Run the Git diff.
	diffs, err := git.Diff(os.Args[1])
	if err != nil {
		errAndExit("An error occurred while getting the Git diff:\n" + err.Error())
	}

	// Create the prompts.
	systemPrompt := util.Must(prompt.CreateSystemPrompt(prompt.SystemInput{
		SuggestionsCount: 5,
	}))
	userPrompt := util.Must(prompt.CreateUserPrompt(prompt.UserInput{
		Diffs: diffs,
	}))

	// Call the OpenAI API.
	client := openai.NewClient(apiKey, openai.Gpt4TurboPreview)
	res := client.CreateCompletion([]openai.Message{
		{
			Role:    openai.RoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.RoleUser,
			Content: userPrompt,
		},
	})

	output := prompt.Output{}
	util.PanicIfError(yaml.Unmarshal([]byte(res.Content), &output))
}

func errAndExit(err string) {
	fmt.Println(err)
	os.Exit(1)
}
