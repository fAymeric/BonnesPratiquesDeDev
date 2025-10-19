package internal

import (
	"math/rand"
	"time"
)

func GameInit(filePath string) Game {
	Word = ChooseWord(filePath)
	GuessedWord = make([]string, len(Word))
	for i := 0; i < len(Word); i++ {
		GuessedWord[i] = "_"
	}
	GuessedLetter = []string{}
	RevealLetters()

	TryAttempt = 10
	return Game{
		GuessedWord: GuessedWord,
		WordSelect:  wordSelect,
		Try:         TryAttempt,
		Alphabet:    Alphabet,
	}
}
func RevealLetters() {
	rand.Seed(time.Now().UnixNano())
	letterRevealed := (len(wordSelect) / 2) - 1

	for i := 0; i < letterRevealed; i++ {
		pos := rand.Intn(len(wordSelect))
		letterToReveal := string(wordSelect[pos])

		for j := 0; j < len(wordSelect); j++ {
			if string(wordSelect[j]) == letterToReveal && GuessedWord[j] == "_" {
				GuessedWord[j] = letterToReveal
				if !contains(GuessedLetter, letterToReveal) {
					GuessedLetter = append(GuessedLetter, letterToReveal)
				}
			}
		}
	}
}
