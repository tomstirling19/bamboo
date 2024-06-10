// openai_service.go
// Handles interactions with OpenAI API. Generic and primary
// purpose is to send requests to OpenAI with basic parsing.

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bamboo/internal/app/models"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

type OpenAIService struct {
	APIKey string
}

func NewOpenAIService(apiKey string) *OpenAIService {
	return &OpenAIService{APIKey: apiKey}
}

func (s *OpenAIService) GetResponse(prompt string) (string, error) {
	requestPayload := models.OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []models.Message{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 150,
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error with response: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var openAIResp models.OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}
