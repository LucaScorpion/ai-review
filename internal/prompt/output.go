package prompt

type Output struct {
	Review          Review           `yaml:"review"`
	CodeSuggestions []CodeSuggestion `yaml:"code_suggestions"`
}

type Review struct {
	PossibleIssues   string `yaml:"possible_issues"`
	SecurityConcerns string `yaml:"security_concerns"`
}

type CodeSuggestion struct {
	RelevantFile string `yaml:"relevant_file"`
	Language     string `yaml:"language"`
	Suggestion   string `yaml:"suggestion"`
	RelevantLine string `yaml:"relevant_line"`
}
