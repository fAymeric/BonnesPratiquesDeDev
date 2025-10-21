package gameLogic

import (
	"fmt"
	"hangman_web/internal/shared"
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
	for _, j := range shared.Word {
		if !contains(shared.GuessedLetter, string(j)) {
			return false
		}
	}
	return true
}

func CheckLetterWord(letter string) {
	for i := 0; i < len(shared.Word); i++ {
		if letter == string(shared.Word[i]) {
			shared.FoundLetter = true
			shared.GuessedWord[i] = letter
		}
	}
}

func CheckLetterAlreadySay(letter string) bool {
	if !contains(shared.GuessedLetter, letter) { // add the letter to guessedletters array
		return true
	}
	fmt.Println("you already guessed that letter")
	return false
}

func UpdateScore(user string) {
	updated := false
	for i, scoreEntry := range shared.Scores {
		if scoreEntry.Pseudo == user {
			shared.Scores[i].Scoreperso = shared.Score
			updated = true
			break
		}
	}
	if !updated && user != "" {
		shared.Scores = append(shared.Scores, shared.ScoreEntry{Pseudo: user, Scoreperso: shared.Score})
	}
}
