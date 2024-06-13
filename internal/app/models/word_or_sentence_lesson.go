package models

type WordOrSentenceLesson struct {
    LessonText        []string `json:"lessonText"`
    EnglishText       []string `json:"englishText"`
    LessonSyllables   []string `json:"lessonSyllables"`
    PhoneticSpellings []string `json:"phoneticSpellings"`
}

func (w *WordOrSentenceLesson) GetContent() [][]string {
    return [][]string{
        w.LessonText,
        w.EnglishText,
        w.LessonSyllables,
        w.PhoneticSpellings,
    }
}
