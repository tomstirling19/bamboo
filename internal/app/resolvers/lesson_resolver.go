package resolvers

import (
	"bamboo/internal/app/models"
	"bamboo/internal/app/services"
	"context"
	"log"
)

type LessonResolver struct {
	LessonService *services.LessonService
    OpenAIService *services.OpenAIService
    Data *models.Lesson
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

func (r *LessonResolver) Description() string {
    return r.Data.Description
}

func (r *LessonResolver) Content() []*LessonContentResolver {
    var content []*LessonContentResolver
    for i := range r.Data.Content {
        content = append(content, &LessonContentResolver{Data: &r.Data.Content[i]})
    }
    return content
}

func (r *LessonResolver) GetLesson(ctx context.Context, req models.LessonRequest) (*LessonResolver, error) {
    request := &models.LessonRequest{
		BaseLesson: models.BaseLesson{
			LessonType: req.LessonType,
			Language:   req.Language,
			Level:      req.Level,
            Topic:      req.Topic,
		},
	}

    prompt := r.LessonService.CreatePrompt(request)
    res, err := r.OpenAIService.GetResponseJSON(prompt)
    if err != nil {
        log.Printf("Error getting response: %v", err)
        return nil, err
    }

    lesson, err := r.LessonService.GetLessonContent(res)
    if err != nil {
        log.Printf("Error parsing response: %v", err)
        return nil, err
    }

    return &LessonResolver{Data: lesson}, nil
}
