package openai

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
)
