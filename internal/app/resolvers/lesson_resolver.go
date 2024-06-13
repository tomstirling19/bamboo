package resolvers

import (
	"bamboo/internal/app/models"
	"bamboo/internal/app/services"
	"context"
	"fmt"
	"log"
)

type LessonResolver struct {
    LessonService *services.LessonService
    OpenAIService *services.OpenAIService
    Data          *models.Lesson
}

func (r *LessonResolver) LessonType() string {
    return r.Data.LessonType
}

func (r *LessonResolver) Language() string {
    return r.Data.Language
}

func (r *LessonResolver) Level() string {
    return r.Data.Level
}

func (r *LessonResolver) Topic() *string {
    return r.Data.Topic
}

func (r *LessonResolver) Description() string {
    return r.Data.Description
}

func (r *LessonResolver) Content() []*LessonContentResolver {
    var contentResolvers []*LessonContentResolver
    for _, content := range r.Data.Content {
        contentResolvers = append(contentResolvers, &LessonContentResolver{Data: content})
    }
    return contentResolvers
}

func (r *LessonResolver) GetLesson(ctx context.Context, req models.LessonRequest) (*LessonResolver, error) {
    prompt := r.LessonService.CreateLessonPrompt(&req)
    res, err := r.OpenAIService.GetResponseJSON(prompt)
    if err != nil {
        log.Printf("Error getting response for lesson type %s: %v", req.LessonType, err)
        return nil, err
    }

    lesson, err := r.LessonService.GetLessonContent(res)
    if err != nil {
        log.Printf("Error parsing response: %v", err)
        return nil, err
    }

    return &LessonResolver{Data: lesson}, nil
}

func (r *LessonContentResolver) ToAlphabetLesson() (*AlphabetLessonResolver, bool) {
    lesson, ok := r.Data.(*models.AlphabetLesson)
    if !ok {
        fmt.Printf("Type mismatch: expected *models.AlphabetLesson, got %T\n", r.Data)
        return nil, false
    }
    return &AlphabetLessonResolver{Data: lesson}, true
}

func (r *LessonContentResolver) ToWordOrSentenceLesson() (*WordOrSentenceLessonResolver, bool) {
    lesson, ok := r.Data.(*models.WordOrSentenceLesson)
    if !ok {
        fmt.Printf("Type mismatch: expected *models.WordOrSentenceLesson, got %T\n", r.Data)
        return nil, false
    }
    return &WordOrSentenceLessonResolver{Data: lesson}, true
}
