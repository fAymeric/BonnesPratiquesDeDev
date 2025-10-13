package internal

import (
	"fmt"
	"strings"
)

func SelecCharac(letter string) Game {
	//fmt.Println("test de lettre : ", letter)
	foundLetter = false
	letter = strings.ToLower(letter)
	if !CheckLetter(letter) || !CheckLetterAlreadySay(letter) {
		return Game{}
	} else {
		if len(letter) <= 1 {
			fmt.Println("test1 de lettre : ", letter)
			CheckLetterWord(letter)
			if foundLetter == false {
				fmt.Println("Your caracter is not in the word")
				LooseTry()
			} else {
				fmt.Println("Your caracter is in the word")
			}
			guessedLetter = append(guessedLetter, letter)
		} else if letter == wordSelect {
			foundLetter = true

		} else {
			LooseTry()
			LooseTry()
		}

	}
	return Game{
		GuessedLetter: guessedLetter,
	}
}
func GetGuessedWord() []string {
	return guessedWord
}
func GetGuessedLetter() []string {
	return guessedLetter
}
func GetTryAttempt() int {
	return tryAttempt
}

func LooseTry() {
	tryAttempt--
}
