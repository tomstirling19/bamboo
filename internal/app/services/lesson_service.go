package services

import (
	"bamboo/internal/app/models"
	"encoding/json"
	"errors"
	"fmt"
)

type LessonService struct {
	OpenAIService *OpenAIService
}

func NewLessonService(openAIService *OpenAIService) *LessonService {
	return &LessonService{OpenAIService: openAIService}
}


func (s *LessonService) GetAlphabetLesson(req *models.LessonRequest) (*models.AlphabetLesson, error) {
    prompt := s.CreateAlphabetLessonPrompt(req)
    res, err := s.OpenAIService.GetJSONResponse(prompt)
    if err != nil {
        return nil, err
    }

    if len(res.Choices) == 0 {
        return nil, errors.New("no choices available in the response")
    }

    contentJson := res.Choices[0].Message.Content.(string)
    var lesson models.AlphabetLesson
    err = json.Unmarshal([]byte(contentJson), &lesson)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON into AlphabetLesson: %v", err)
    }
    return &lesson, nil
}


func (s *LessonService) GetWordOrSentenceLesson(req *models.LessonRequest) (*models.WordOrSentenceLesson, error) {
    var prompt string

    switch req.LessonType {
    case "word":
        prompt = s.CreateWordLessonPrompt(req)
    case "sentence":
        prompt = s.CreateSentenceLessonPrompt(req)
    default:
        return nil, fmt.Errorf("unknown lesson type: %s", req.LessonType)
    }

    res, err := s.OpenAIService.GetJSONResponse(prompt)
    if err != nil {
        return nil, err
    }

    if len(res.Choices) == 0 {
        return nil, errors.New("no choices available in the response")
    }

    contentJson := res.Choices[0].Message.Content.(string)
    var lesson models.WordOrSentenceLesson
    err = json.Unmarshal([]byte(contentJson), &lesson)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON into WordOrSentenceLesson: %v", err)
    }
    return &lesson, nil
}
