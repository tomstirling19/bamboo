package resolvers

import (
	"bamboo/internal/app/models"
	"bamboo/internal/app/services"
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/graphql-go"
)

type LessonResolver struct {
    LessonService *services.LessonService
    OpenAIService *services.OpenAIService
}

func NewLessonSchema(resolver *LessonResolver) (*graphql.Schema, error) {
    schema := `
        schema {
            query: Query
        }
        type Query {
            getLesson(language: String!, level: String!, topic: String!): LessonResponse!
        }
        type LessonResponse {
            language: String!
            level: String!
            title: String!
            description: String!
            content: [TranslationPair!]!
        }
        type TranslationPair {
            lessonText: String!
            englishText: String!
        }
    `

    parsedSchema, err := graphql.ParseSchema(schema, resolver)
    if err != nil {
        return nil, fmt.Errorf("failed to parse schema: %w", err)
    }

    return parsedSchema, nil
}

func (r *LessonResolver) GetLesson(ctx context.Context, req models.LessonRequest) (*models.LessonResolver, error) {
    request := &models.LessonRequest{
        Language: req.Language,
        Level:    req.Level,
        Topic:    req.Topic,
    }

    prompt := r.LessonService.CreatePrompt(request)
    res, err := r.OpenAIService.GetResponseJSON(prompt)
    if err != nil {
        log.Printf("Error getting response: %v", err)
        return nil, err
    }

    lessonResponse, err := r.LessonService.GetLessonContent(res)
    if err != nil {
        log.Printf("Error parsing response: %v", err)
        return nil, err
    }

    return &models.LessonResolver{Data: lessonResponse}, nil
}
