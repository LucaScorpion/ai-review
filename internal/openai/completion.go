package openai

func (c *Client) CreateCompletion(messages []Message) Message {
	res := &completionResponse{}
	c.request(
		"POST",
		"/v1/chat/completions",
		completionRequest{
			Model:    c.model,
			Messages: messages,
		},
		res,
	)

	return res.Choices[0].Message
}

type completionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type completionResponse struct {
	Choices []completionChoice `json:"choices"`
}

type completionChoice struct {
	Message Message `json:"message"`
}
