package resolvers

import "bamboo/internal/app/models"

type AlphabetLessonResolver struct {
    Data *models.AlphabetLesson
}

type AlphabetContentResolver struct {
    Data *models.AlphabetContent
}

func (r *AlphabetLessonResolver) LessonType() string {
    return r.Data.LessonType
}

func (r *AlphabetLessonResolver) Language() string {
    return r.Data.Language
}

func (r *AlphabetLessonResolver) Level() string {
    return r.Data.Level
}

func (r *AlphabetLessonResolver) Description() *string {
    return r.Data.Description
}

func (r *AlphabetLessonResolver) Content() []*AlphabetContentResolver {
    resolvers := make([]*AlphabetContentResolver, len(r.Data.Content))
    for i, content := range r.Data.Content {
        resolvers[i] = &AlphabetContentResolver{Data: &content}
    }
    return resolvers
}

func (r *AlphabetContentResolver) AlphabetCharacter() []string {
    return r.Data.AlphabetCharacter
}

func (r *AlphabetContentResolver) Phoneme() []string {
    return r.Data.Phoneme
}
