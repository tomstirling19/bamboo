package resolvers

import "bamboo/internal/app/models"

type LessonContentResolver struct {
    Data *models.LessonContent
}

func (r *LessonContentResolver) LessonText() []string {
    return r.Data.LessonText
}

func (r *LessonContentResolver) EnglishText() []string {
    return r.Data.EnglishText
}

func (r *LessonContentResolver) LessonSyllables() []string {
    return r.Data.LessonSyllables
}

func (r *LessonContentResolver) PhoneticSpellings() []string {
    return r.Data.PhoneticSpellings
}
