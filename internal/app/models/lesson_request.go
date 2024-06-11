package models

type LessonRequest struct {
    Language string `json:"language"`
    Level    string `json:"level"`
    Topic    string `json:"topic"`
}
