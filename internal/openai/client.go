package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultEndpoint = "https://api.openai.com/v1/chat/completions"
const defaultModel = "gpt-3.5-turbo"

type OpenAIClient struct {
	apiKey     string
	maxTokens  int
	httpClient *http.Client
	endpoint   string
	model      string
}

func NewOpenAIClient(apiKey string, maxTokens int, client *http.Client, endpoint string, model string) *OpenAIClient {
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	if model == "" {
		model = defaultModel
	}
	return &OpenAIClient{
		apiKey:     apiKey,
		maxTokens:  maxTokens,
		httpClient: client,
		endpoint:   endpoint,
		model:      model,
	}
}

func (aiClient *OpenAIClient) GetAnswer(prompt string) (string, error) {
	data := Request{
		Messages:  []Message{{Role: "user", Content: prompt}},
		Model:     aiClient.model,
		MaxTokens: aiClient.maxTokens,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", aiClient.endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+aiClient.apiKey)

	resp, err := aiClient.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("received response code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
