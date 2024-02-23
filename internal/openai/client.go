package openai

import (
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
	jsonBody := must(json.Marshal(body))

	req := must(http.NewRequest(method, baseUrl+path, bytes.NewReader(jsonBody)))
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res := must(http.DefaultClient.Do(req))
	resBody := must(io.ReadAll(res.Body))

	if res.StatusCode != 200 {
		fmt.Printf("Received an API error (%d)\n", res.StatusCode)
		panic(string(resBody))
	}

	panicIfError(json.Unmarshal(resBody, resVal))
}

func must[T any](val T, err error) T {
	panicIfError(err)
	return val
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
