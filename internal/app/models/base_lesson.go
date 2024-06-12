package models

type BaseLesson struct {
	LessonType  string `json:"lessonType"`
    Language    string `json:"language"`
    Level       string `json:"level"`
    Topic       string `json:"topic"`
}
