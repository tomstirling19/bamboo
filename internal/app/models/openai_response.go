package models

type Choice struct {
	Message Message `json:"message"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}
