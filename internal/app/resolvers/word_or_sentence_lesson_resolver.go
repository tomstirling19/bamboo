package resolvers

import "bamboo/internal/app/models"

type WordOrSentenceLessonResolver struct {
    Data *models.WordOrSentenceLesson
}

func (r *WordOrSentenceLessonResolver) LessonText() []string {
    return r.Data.LessonText
}

func (r *WordOrSentenceLessonResolver) EnglishText() []string {
    return r.Data.EnglishText
}

func (r *WordOrSentenceLessonResolver) LessonSyllables() []string {
    return r.Data.LessonSyllables
}

func (r *WordOrSentenceLessonResolver) PhoneticSpellings() []string {
    return r.Data.PhoneticSpellings
}
