package internal

import (
	"html/template"
	"net/http"
)

func EnEasyPage(w http.ResponseWriter, r *http.Request) { //calls GamePage func with easy mode
	EnGamePage(w, r, "easy")
}

func EnMediumPage(w http.ResponseWriter, r *http.Request) { //calls GamePage func with medium mode
	EnGamePage(w, r, "medium")
}

func EnHardPage(w http.ResponseWriter, r *http.Request) { //calls GamePage func with hard
	EnGamePage(w, r, "hard")
}

func EnGamePage(w http.ResponseWriter, r *http.Request, difficulty string) {
	Tg := template.Must(template.ParseFiles("web/template/En/GamePageEn.html"))

	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Set word file based on difficulty
	var wordFile string
	switch difficulty {
	case "easy":
		wordFile = "internal/data/english/words.txt"
	case "medium":
		wordFile = "internal/data/english/words2.txt"
	case "hard":
		wordFile = "internal/data/english/words3.txt"
	default:
		http.Error(w, "Invalid difficulty", http.StatusBadRequest)
		return
	}

	// Initialize game if not already done
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit(wordFile))
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
			http.Redirect(w, r, "/EnLoose", http.StatusFound)
			return
		}
		if CheckVictory() || currentGame.WordSelect == guessWord {
			http.Redirect(w, r, "/EnWin", http.StatusFound)
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

	Tg.Execute(w, currentGame)
}
