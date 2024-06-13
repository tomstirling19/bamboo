package resolvers

import "bamboo/internal/app/models"

type AlphabetLessonResolver struct {
    Data *models.AlphabetLesson
}

func (r *AlphabetLessonResolver) AlphabetCharacter() []string {
    return r.Data.AlphabetCharacter
}

func (r *AlphabetLessonResolver) Phoneme() []string {
    return r.Data.Phoneme
}
