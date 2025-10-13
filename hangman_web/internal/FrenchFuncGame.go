package internal

import (
	"html/template"
	"net/http"
)

func FrEasyPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/FrEasy.html"))
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/french/words.txt"))
	}
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
			http.Redirect(w, r, "/FrLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
			http.Redirect(w, r, "/FrWin", http.StatusFound)
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
func FrMediumPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/FrMedium.html"))
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/french/words2.txt"))
	}
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
			http.Redirect(w, r, "/FrLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
			http.Redirect(w, r, "/FrWin", http.StatusFound)
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
func FrHardPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/FrHard.html"))
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/french/words3.txt"))
	}
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
			http.Redirect(w, r, "/FrLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
			http.Redirect(w, r, "/FrWin", http.StatusFound)
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
