package models

type AlphabetLesson struct {
    AlphabetCharacter []string `json:"alphabetCharacter"`
    Phoneme           []string `json:"phoneme"`
}

func (a *AlphabetLesson) GetContent() [][]string {
    return [][]string{
        a.AlphabetCharacter,
        a.Phoneme,
    }
}