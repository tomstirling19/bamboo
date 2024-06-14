package resolvers

import "bamboo/internal/app/models"

type WordOrSentenceLessonResolver struct {
    Data *models.WordOrSentenceLesson
}

type WordOrSentenceContentResolver struct {
    Data *models.WordOrSentenceContent
}

func (r *WordOrSentenceLessonResolver) LessonType() string {
    return r.Data.LessonType
}

func (r *WordOrSentenceLessonResolver) Language() string {
    return r.Data.Language
}

func (r *WordOrSentenceLessonResolver) Level() string {
    return r.Data.Level
}

func (r *WordOrSentenceLessonResolver) Topic() *string {
    return r.Data.Topic
}

func (r *WordOrSentenceLessonResolver) Description() *string {
    return r.Data.Description
}

func (r *WordOrSentenceLessonResolver) Content() []*WordOrSentenceContentResolver {
    resolvers := make([]*WordOrSentenceContentResolver, len(r.Data.Content))
    for i, content := range r.Data.Content {
        resolvers[i] = &WordOrSentenceContentResolver{Data: &content}
    }
    return resolvers
}

func (r *WordOrSentenceContentResolver) LessonText() []string {
    return r.Data.LessonText
}

func (r *WordOrSentenceContentResolver) EnglishText() []string {
    return r.Data.EnglishText
}

func (r *WordOrSentenceContentResolver) LessonSyllables() []string {
    return r.Data.LessonSyllables
}

func (r *WordOrSentenceContentResolver) PhoneticSpellings() []string {
    return r.Data.PhoneticSpellings
}
