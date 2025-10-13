package internal

import (
	"html/template"
	"net/http"
)

func DeutschEasyPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Deutsch/DeutschEasy.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/deutsch/words.txt"))
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
			http.Redirect(w, r, "/DeutschLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
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

	Tg.Execute(w, currentGame)
}

func DeutschMediumPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Deutsch/DeutschMedium.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/deutsch/words2.txt"))
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
			http.Redirect(w, r, "/DeutschLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
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

	Tg.Execute(w, currentGame)
}
func DeutschHardPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Deutsch/DeutschHard.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/deutsch/words3.txt"))
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
			http.Redirect(w, r, "/DeutschLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
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

	Tg.Execute(w, currentGame)
}
