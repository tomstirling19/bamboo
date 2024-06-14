package models

type WordOrSentenceLesson struct {
    BaseLesson
    Content []WordOrSentenceContent `json:"content"`
}

type WordOrSentenceContent struct {
    LessonText        []string `json:"lessonText"`
    EnglishText       []string `json:"englishText"`
    LessonSyllables   []string `json:"lessonSyllables"`
    PhoneticSpellings []string `json:"phoneticSpellings"`
}
