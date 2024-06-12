package models

type Lesson struct {
    BaseLesson
    Description string          `json:"description"`
    Content     []LessonContent `json:"content"`
}
