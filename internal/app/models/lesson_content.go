package models

type LessonContent struct {
    LessonText  	  []string `json:"lessonText"`
    EnglishText 	  []string `json:"englishText"`
    LessonSyllables   []string `json:"lessonSyllables"`
    PhoneticSpellings []string `json:"phoneticSpellings"`
}
