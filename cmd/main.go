package main

import (
	"ai-reviewer/internal/ansi"
	"ai-reviewer/internal/git"
	"ai-reviewer/internal/openai"
	"ai-reviewer/internal/prompt"
	"ai-reviewer/internal/util"
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

//go:embed banner.txt
var banner string

func main() {
	// Check if the OpenAI API key is set.
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		errAndExit("No OPENAI_API_KEY found in environment.")
	}

	// Check if we have a target branch.
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

	fmt.Println(banner)

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

	response := strings.TrimSuffix(res.Content, "```")

	// Parse the yaml, check for errors.
	output := prompt.Output{}
	err = yaml.Unmarshal([]byte(response), &output)
	if err != nil {
		fmt.Println("Failed to unmarshal yaml:", err)
		fmt.Println(response)
		os.Exit(1)
	}

	// Display the output.

	fmt.Println(ansi.Bold() + "Possible Issues:" + ansi.Reset())
	fmt.Println(output.Review.PossibleIssues)

	fmt.Println(ansi.Bold() + "Security Concerns:" + ansi.Reset())
	fmt.Println(output.Review.SecurityConcerns)

	fmt.Println(ansi.Bold() + "Suggestions:" + ansi.Reset() + "\n")
	for _, suggestion := range output.CodeSuggestions {
		fmt.Println(ansi.Bold() + strings.TrimSpace(suggestion.RelevantFile) + ansi.Reset())
		fmt.Println("```\n" + strings.TrimSpace(suggestion.RelevantLine) + "\n```")
		fmt.Println(suggestion.Suggestion)
	}
}

func errAndExit(err string) {
	fmt.Println(err)
	os.Exit(1)
}
