package services

import (
	"bamboo/internal/app/models"
	"encoding/json"
	"fmt"
)

type LessonService struct{}

func NewLessonService() *LessonService {
	return &LessonService{}
}

func (s *LessonService) CreatePrompt(request *models.LessonRequest) string {
	return fmt.Sprintf(
		`Generate a list of %d translations in %s for a %s lesson on the topic "%s". 
        Provide translations for both words and sentences in %s and their English equivalents. 
        Return a JSON object with translation pairs in the following format:
        {"language": "%s","level": "%s","title": "","description": "","content": [{"lessonText": "","englishText": ""}]}`,
		8, request.Language, request.Level, request.Topic, request.Language, request.Language, request.Level)
}

func (s *LessonService) GetLessonContent(res *models.OpenAIResponse) (*models.Lesson, error) {
	if len(res.Choices) == 0 || res.Choices[0].Message.Content == "" {
		return nil, fmt.Errorf("no content found in response")
	}

	contentJSON := res.Choices[0].Message.Content.(string)

	var lessonData models.Lesson
	if err := json.Unmarshal([]byte(contentJSON), &lessonData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal lesson content: %v", err)
	}

	return &lessonData, nil
}
