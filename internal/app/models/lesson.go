package models

type Lesson struct {
    Language    string            `json:"language"`
    Level       string            `json:"level"`
    Title       string            `json:"title"`
    Description string            `json:"description"`
    Content     []TranslationData `json:"content"`
}

type TranslationData struct {
    LessonText  string `json:"lessonText"`
    EnglishText string `json:"englishText"`
}

type LessonResolver struct {
    Data *Lesson
}

type TranslationDataResolver struct {
    Data *TranslationData
}

func (r *LessonResolver) Language() string {
    return r.Data.Language
}

func (r *LessonResolver) Level() string {
    return r.Data.Level
}

func (r *LessonResolver) Title() string {
    return r.Data.Title
}

func (r *LessonResolver) Description() string {
    return r.Data.Description
}

func (r *LessonResolver) Content() []*TranslationDataResolver {
    var content []*TranslationDataResolver
    for i := range r.Data.Content {
        content = append(content, &TranslationDataResolver{Data: &r.Data.Content[i]})
    }
    return content
}

func (r *TranslationDataResolver) LessonText() string {
    return r.Data.LessonText
}

func (r *TranslationDataResolver) EnglishText() string {
    return r.Data.EnglishText
}
