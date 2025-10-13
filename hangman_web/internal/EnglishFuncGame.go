package internal

import (
	"html/template"
	"net/http"
)

func EnEasyPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/En/EnEasy.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/english/words.txt"))
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
			http.Redirect(w, r, "/EnLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
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
func EnMediumPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/En/EnMedium.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/english/words2.txt"))
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
			http.Redirect(w, r, "/EnLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
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
func EnHardPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/En/EnHard.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if currentGame.WordSelect == "" {
		currentGame = Game(GameInit("internal/data/english/words3.txt"))
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
			http.Redirect(w, r, "/EnLoose", http.StatusFound)
			return
		}
		if CheckVictory() == true || currentGame.WordSelect == guessWord {
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
