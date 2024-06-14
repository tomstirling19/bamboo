package resolvers

import (
	"bamboo/internal/app/models"
	"bamboo/internal/app/services"
	"context"
)

type LessonResolver struct {
    LessonService *services.LessonService
}

func (r *LessonResolver) GetAlphabetLesson(ctx context.Context, args struct{ Language, Level string }) (*AlphabetLessonResolver, error) {
    request := &models.LessonRequest{
        BaseLesson: models.BaseLesson{
            LessonType:  "alphabet",
            Language:    args.Language,
            Level:       args.Level,
            Topic:       nil,
            Description: nil,
        },
    }

    lesson, err := r.LessonService.GetAlphabetLesson(request)
    if err != nil {
        return nil, err
    }

    return &AlphabetLessonResolver{Data: lesson}, nil
}

func (r *LessonResolver) GetWordOrSentenceLesson(ctx context.Context, args struct{ LessonType, Language, Level, Topic string }) (*WordOrSentenceLessonResolver, error) {
    request := &models.LessonRequest{
        BaseLesson: models.BaseLesson{
            LessonType:  args.LessonType,
            Language:    args.Language,
            Level:       args.Level,
            Topic:       &args.Topic,
            Description: nil,
        },
    }

    lesson, err := r.LessonService.GetWordOrSentenceLesson(request)
    if err != nil {
        return nil, err
    }

    return &WordOrSentenceLessonResolver{Data: lesson}, nil
}
