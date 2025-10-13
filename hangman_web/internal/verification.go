package internal

import (
	"fmt"
)

func contains(slice []string, item string) bool { // use for checking if a letter is already in an array
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func CheckLetter(letter string) bool {
	if len(letter) <= 1 {
		if !(letter >= "a" && letter <= "z") {
			//fmt.Println("Only a letter or a full word is needed")
			//fmt.Println("Your caracter is invalid")
			return false
		}
	}
	return true
}

func CheckVictory() bool { // looking for all the letters and compare to wordselect
	for _, j := range word {
		if !contains(guessedLetter, string(j)) {
			return false
		}
	}
	return true
}

func CheckLetterWord(letter string) {
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) {
			foundLetter = true
			guessedWord[i] = letter
		}
	}
}

func CheckLetterAlreadySay(letter string) bool {
	if !contains(guessedLetter, letter) { // add the letter to guessedletters array
		return true
	}
	fmt.Println("you already guessed that letter")
	return false
}
