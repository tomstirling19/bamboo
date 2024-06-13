package services

import (
	"bamboo/internal/app/models"
	"bamboo/internal/app/utils"
	"bamboo/internal/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

type OpenAIService struct {
	Config *config.OpenAIConfig
}

func NewOpenAIService(config *config.OpenAIConfig) *OpenAIService {
	return &OpenAIService{Config: config}
}

func (s *OpenAIService) GetResponseJSON(prompt string) (*models.OpenAIResponse, error) {
	requestPayload, err := createRequestPayload(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to create request payload: %v", err)
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %v", err)
	}

	req, err := createHTTPRequest(s.Config.APIKey, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	responseBody, err := executeHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	openAIResponse, err := parseResponseBody(responseBody)
	if err != nil {
		return nil, err
	}

	log.Printf("OpenAI Response: %s", utils.ToJSONString(openAIResponse))
	return openAIResponse, nil
}

func createRequestPayload(prompt string) (*models.OpenAIRequest, error) {
	return &models.OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []models.Message{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1000,
	}, nil
}

func createHTTPRequest(apiKey string, requestBody []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	return req, nil
}

func executeHTTPRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error with response: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return io.ReadAll(resp.Body)
}

func parseResponseBody(responseBody []byte) (*models.OpenAIResponse, error) {
	var openAIResponse models.OpenAIResponse
	err := json.Unmarshal(responseBody, &openAIResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return &openAIResponse, nil
}
