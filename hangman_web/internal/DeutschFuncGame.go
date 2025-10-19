package internal

import (
	"html/template"
	"net/http"
)

func DeutschEasyPage(w http.ResponseWriter, r *http.Request) {
	DeutschGamePage(w, r, "easy")
}

func DeutschMediumPage(w http.ResponseWriter, r *http.Request) {
	DeutschGamePage(w, r, "medium")
}

func DeutschHardPage(w http.ResponseWriter, r *http.Request) {
	DeutschGamePage(w, r, "hard")
}

func DeutschGamePage(w http.ResponseWriter, r *http.Request, difficulty string) {
	Tg := template.Must(template.ParseFiles("web/template/Deutsch/GamePageDeutsch.html"))

	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Set word file based on difficulty
	var wordFile string
	switch difficulty {
	case "easy":
		wordFile = "internal/data/deutsch/words.txt"
	case "medium":
		wordFile = "internal/data/deutsch/words2.txt"
	case "hard":
		wordFile = "internal/data/deutsch/words3.txt"
	default:
		http.Error(w, "Invalid difficulty", http.StatusBadRequest)
		return
	}

	// Initialize game if not already done
	if currentGame.WordSelect == "" {
		currentGame = GameInit(wordFile)
	}

	// Handle form POST
	if r.Method == http.MethodPost {
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")

		SelecCharac(guess)
		SelecCharac(guessWord)

		currentGame.GuessedWord = GetGuessedWord()
		currentGame.GuessedLetter = GetGuessedLetter()
		currentGame.Try = GetTryAttempt()
		currentGame.Alphabet = alphabet

		if GetTryAttempt() == 0 {
			http.Redirect(w, r, "/DeutschLoose", http.StatusFound)
			return
		}
		if CheckVictory() || currentGame.WordSelect == guessWord {
			http.Redirect(w, r, "/DeutschWin", http.StatusFound)
			Score++
			updated := false
			for i, scoreEntry := range scores {
				if scoreEntry.Pseudo == user {
					scores[i].Scoreperso = Score
					updated = true
					break
				}
			}
			if !updated && user != "" {
				scores = append(scores, ScoreEntry{Pseudo: user, Scoreperso: Score})
			}
			return
		}
	}

	err := Tg.Execute(w, currentGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
