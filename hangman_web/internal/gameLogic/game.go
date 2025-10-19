package gameLogic

import (
	"hangman_web/internal/shared"
	"math/rand"
	"time"
)

func GameInit(filePath string) shared.Game {
	shared.Word = ChooseWord(filePath)
	shared.GuessedWord = make([]string, len(shared.Word))
	for i := 0; i < len(shared.Word); i++ {
		shared.GuessedWord[i] = "_"
	}
	shared.GuessedLetter = []string{}
	RevealLetters()

	shared.TryAttempt = 10
	return shared.Game{
		GuessedWord: shared.GuessedWord,
		WordSelect:  wordSelect,
		Try:         shared.TryAttempt,
		Alphabet:    shared.Alphabet,
	}
}
func RevealLetters() {
	rand.Seed(time.Now().UnixNano())
	letterRevealed := (len(wordSelect) / 2) - 1

	for i := 0; i < letterRevealed; i++ {
		pos := rand.Intn(len(wordSelect))
		letterToReveal := string(wordSelect[pos])

		for j := 0; j < len(wordSelect); j++ {
			if string(wordSelect[j]) == letterToReveal && shared.GuessedWord[j] == "_" {
				shared.GuessedWord[j] = letterToReveal
				if !contains(shared.GuessedLetter, letterToReveal) {
					shared.GuessedLetter = append(shared.GuessedLetter, letterToReveal)
				}
			}
		}
	}
}
