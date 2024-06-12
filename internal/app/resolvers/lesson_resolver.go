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

func NewLessonSchema(resolver *LessonResolver) (*graphql.Schema, error) {
    schema := `
        schema {
            query: Query
        }
        type Query {
            getLesson(lessonType: String!, language: String!, level: String!, topic: String!): Lesson!
        }
        type Lesson {
            lessonType: String!
            language: String!
            level: String!
            description: String!
            content: [LessonContent!]!
        }
        type LessonContent {
            lessonText: [String!]!
            englishText: [String!]!
			lessonSyllables: [String!]!
            phoneticSpellings: [String!]!
        }
    `

    parsedSchema, err := graphql.ParseSchema(schema, resolver)
    if err != nil {
        return nil, fmt.Errorf("failed to parse schema: %w", err)
    }

    return parsedSchema, nil
}
