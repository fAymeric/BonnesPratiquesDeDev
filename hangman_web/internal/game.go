package internal

import (
	"math/rand"
	"time"
)

func GameInit(filePath string) Game {
	word = ChooseWord(filePath)
	guessedWord = make([]string, len(word))
	for i := 0; i < len(word); i++ {
		guessedWord[i] = "_"
	}
	guessedLetter = []string{}
	RevealLetters()

	tryAttempt = 10
	return Game{
		GuessedWord: guessedWord,
		WordSelect:  wordSelect,
		Try:         tryAttempt,
	}
}
func RevealLetters() {
	rand.Seed(time.Now().UnixNano())
	letterRevealed := (len(wordSelect) / 2) - 1

	for i := 0; i < letterRevealed; i++ {
		pos := rand.Intn(len(wordSelect))
		letterToReveal := string(wordSelect[pos])

		for j := 0; j < len(wordSelect); j++ {
			if string(wordSelect[j]) == letterToReveal && guessedWord[j] == "_" {
				guessedWord[j] = letterToReveal
				if !contains(guessedLetter, letterToReveal) {
					guessedLetter = append(guessedLetter, letterToReveal)
				}
			}
		}
	}
}
