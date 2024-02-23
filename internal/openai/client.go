package openai

import (
	"ai-reviewer/internal/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://api.openai.com"

const (
	Gpt4TurboPreview = "gpt-4-turbo-preview"
)

type Client struct {
	apiKey string
	model  string
}

func NewClient(apiKey, model string) *Client {
	return &Client{
		apiKey: apiKey,
		model:  model,
	}
}

func (c *Client) request(method, path string, body, resVal any) {
	jsonBody := util.Must(json.Marshal(body))

	req := util.Must(http.NewRequest(method, baseUrl+path, bytes.NewReader(jsonBody)))
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res := util.Must(http.DefaultClient.Do(req))
	resBody := util.Must(io.ReadAll(res.Body))

	if res.StatusCode != 200 {
		fmt.Printf("Received an API error (%d)\n", res.StatusCode)
		panic(string(resBody))
	}

	util.PanicIfError(json.Unmarshal(resBody, resVal))
}
