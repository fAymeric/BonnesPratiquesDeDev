package gameLogic

import (
	"fmt"
	"hangman_web/internal/shared"
	"strings"
)

func SelecCharac(letter string) shared.Game {
	//fmt.Println("test de lettre : ", letter)
	shared.FoundLetter = false
	letter = strings.ToLower(letter)
	if !CheckLetter(letter) || !CheckLetterAlreadySay(letter) {
		return shared.Game{}
	} else {
		if len(letter) <= 1 {
			fmt.Println("test1 de lettre : ", letter)
			CheckLetterWord(letter)
			if shared.FoundLetter == false {
				fmt.Println("Your caracter is not in the word")
				LooseTry()
			} else {
				fmt.Println("Your caracter is in the word")
			}
			shared.GuessedLetter = append(shared.GuessedLetter, letter)
		} else if letter == wordSelect {
			shared.FoundLetter = true

		} else {
			LooseTry()
			LooseTry()
		}

	}
	return shared.Game{
		GuessedLetter: shared.GuessedLetter,
	}
}
func GetGuessedWord() []string {
	return shared.GuessedWord
}
func GetGuessedLetter() []string {
	return shared.GuessedLetter
}
func GetTryAttempt() int {
	return shared.TryAttempt
}

func LooseTry() {
	shared.TryAttempt--
}
