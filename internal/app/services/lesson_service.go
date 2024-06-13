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

func (s *LessonService) GetLessonContent(res *models.OpenAIResponse) (*models.Lesson, error) {
    contentJSON, err := extractContent(res)
    if err != nil {
        return nil, err
    }

    var raw map[string]json.RawMessage
    if err := json.Unmarshal([]byte(contentJSON), &raw); err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    lesson, err := createBaseLesson(raw)
    if err != nil {
        return nil, err
    }

    lessonContent, err := unmarshalLessonContent(raw["content"], lesson.LessonType)
    if err != nil {
        return nil, err
    }

    lesson.Content = lessonContent
    return lesson, nil
}

func extractContent(res *models.OpenAIResponse) (string, error) {
    if len(res.Choices) == 0 || res.Choices[0].Message.Content == "" {
        return "", fmt.Errorf("no content found in response")
    }
    content, ok := res.Choices[0].Message.Content.(string)
    if !ok {
        return "", fmt.Errorf("invalid content format")
    }
    return content, nil
}

func createBaseLesson(raw map[string]json.RawMessage) (*models.Lesson, error) {
    var lesson models.Lesson

    fields := map[string]interface{}{
        "lessonType": &lesson.LessonType,
        "language":   &lesson.Language,
        "level":      &lesson.Level,
        "description": &lesson.Description,
    }

    for key, field := range fields {
        if err := json.Unmarshal(raw[key], field); err != nil {
            return nil, fmt.Errorf("failed to unmarshal %s: %w", key, err)
        }
    }

    return &lesson, nil
}

func unmarshalLessonContent(contentJSON json.RawMessage, lessonType string) ([]models.LessonContent, error) {
    var contentArray []json.RawMessage
    if err := json.Unmarshal(contentJSON, &contentArray); err != nil {
        return nil, fmt.Errorf("failed to unmarshal content array: %w", err)
    }

    var lessonContent []models.LessonContent

    for _, content := range contentArray {
        lesson, err := unmarshalContentByType(content, lessonType)
        if err != nil {
            return nil, err
        }
        lessonContent = append(lessonContent, lesson)
    }

    return lessonContent, nil
}

func unmarshalContentByType(content json.RawMessage, lessonType string) (models.LessonContent, error) {
    switch lessonType {
    case "alphabet":
        var lesson models.AlphabetLesson
        if err := json.Unmarshal(content, &lesson); err != nil {
            return nil, fmt.Errorf("failed to unmarshal alphabet content: %w", err)
        }
        return &lesson, nil
    case "word", "sentence":
        var lesson models.WordOrSentenceLesson
        if err := json.Unmarshal(content, &lesson); err != nil {
            return nil, fmt.Errorf("failed to unmarshal word/sentence content: %w", err)
        }
        return &lesson, nil
    default:
        return nil, fmt.Errorf("unknown lesson type: %s", lessonType)
    }
}

func (s *LessonService) CreateLessonPrompt(request *models.LessonRequest) string {
    switch request.LessonType {
    case "alphabet":
        return s.createAlphabetLessonPrompt(request)
    case "word":
        return s.createWordLessonPrompt(request)
    case "sentence":
        return s.createSentenceLessonPrompt(request)
    default:
        return ""
    }
}