package prompt

import (
	"ai-reviewer/internal/git"
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed prompts/system.txt
var systemPrompt string

//go:embed prompts/user.txt
var userPrompt string

type SystemInput struct {
	SuggestionsCount int
}

type UserInput struct {
	Diffs []git.FileDiff
}

func CreateSystemPrompt(input SystemInput) (string, error) {
	return executeTemplate(systemPrompt, input)
}

func CreateUserPrompt(input UserInput) (string, error) {
	return executeTemplate(userPrompt, input)
}

func executeTemplate(text string, input any) (string, error) {
	tpl := template.Must(template.New("").Parse(text))
	var buf bytes.Buffer
	err := tpl.Execute(&buf, input)
	return buf.String(), err
}
