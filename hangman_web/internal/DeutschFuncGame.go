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

	if User == "" {
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
	if CurrentGame.WordSelect == "" {
		CurrentGame = GameInit(wordFile)
	}

	// Handle form POST
	if r.Method == http.MethodPost {
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")

		SelecCharac(guess)
		SelecCharac(guessWord)

		CurrentGame.GuessedWord = GetGuessedWord()
		CurrentGame.GuessedLetter = GetGuessedLetter()
		CurrentGame.Try = GetTryAttempt()
		CurrentGame.Alphabet = Alphabet

		if GetTryAttempt() == 0 {
			RenderLoosePage(w, r, "de")
			return
		}
		if CheckVictory() || CurrentGame.WordSelect == guessWord {
			RenderWinPage(w, r, "de")
			Score++
			updated := false
			for i, scoreEntry := range scores {
				if scoreEntry.Pseudo == User {
					scores[i].Scoreperso = Score
					updated = true
					break
				}
			}
			if !updated && User != "" {
				scores = append(scores, ScoreEntry{Pseudo: User, Scoreperso: Score})
			}
			return
		}
	}

	err := Tg.Execute(w, CurrentGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
