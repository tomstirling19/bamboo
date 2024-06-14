package models

type AlphabetLesson struct {
    BaseLesson
    Content []AlphabetContent `json:"content"`
}

type AlphabetContent struct {
    AlphabetCharacter []string `json:"alphabetCharacter"`
    Phoneme           []string `json:"phoneme"`
}
