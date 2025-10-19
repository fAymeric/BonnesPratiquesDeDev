package internal

import (
	"fmt"
	"strings"
)

func SelecCharac(letter string) Game {
	//fmt.Println("test de lettre : ", letter)
	FoundLetter = false
	letter = strings.ToLower(letter)
	if !CheckLetter(letter) || !CheckLetterAlreadySay(letter) {
		return Game{}
	} else {
		if len(letter) <= 1 {
			fmt.Println("test1 de lettre : ", letter)
			CheckLetterWord(letter)
			if FoundLetter == false {
				fmt.Println("Your caracter is not in the word")
				LooseTry()
			} else {
				fmt.Println("Your caracter is in the word")
			}
			GuessedLetter = append(GuessedLetter, letter)
		} else if letter == wordSelect {
			FoundLetter = true

		} else {
			LooseTry()
			LooseTry()
		}

	}
	return Game{
		GuessedLetter: GuessedLetter,
	}
}
func GetGuessedWord() []string {
	return GuessedWord
}
func GetGuessedLetter() []string {
	return GuessedLetter
}
func GetTryAttempt() int {
	return TryAttempt
}

func LooseTry() {
	TryAttempt--
}
