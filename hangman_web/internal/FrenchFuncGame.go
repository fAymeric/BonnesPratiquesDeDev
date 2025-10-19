package internal

import (
	"html/template"
	"net/http"
)

func FrEasyPage(w http.ResponseWriter, r *http.Request) {
	FrGamePage(w, r, "easy")
}
func FrMediumPage(w http.ResponseWriter, r *http.Request) {
	FrGamePage(w, r, "medium")
}
func FrHardPage(w http.ResponseWriter, r *http.Request) {
	FrGamePage(w, r, "hard")
}

func FrGamePage(w http.ResponseWriter, r *http.Request, difficulty string) {
	Tg := template.Must(template.ParseFiles("web/template/Game.html"))

	if User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Set word file based on difficulty
	var wordFile string
	switch difficulty {
	case "easy":
		wordFile = "internal/data/french/words.txt"
	case "medium":
		wordFile = "internal/data/french/words2.txt"
	case "hard":
		wordFile = "internal/data/french/words3.txt"
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
			RenderLoosePage(w, r, "fr")
			return
		}
		if CheckVictory() || CurrentGame.WordSelect == guessWord {
			RenderWinPage(w, r, "fr")
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
