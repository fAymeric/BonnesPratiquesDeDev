package internal

import (
	Hangman "HangmanWeb/OriginalHangman/pkg"
	"html/template"
	"net/http"
)

func FrEasyPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/FR/FrEasy.html"))
	if currentGame.WordSelect == "" {
		currentGame = Game(Hangman.GameInit("OriginalHangman/data/french/words.txt"))
	}
	if r.Method == http.MethodPost {
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")
		Hangman.SelecCharac(guess)
		Hangman.SelecCharac(guessWord)
		currentGame.GuessedWord = Hangman.GetGuessedWord()
		currentGame.GuessedLetter = Hangman.GetGuessedLetter()
		currentGame.Try = Hangman.GetTryAttempt()
		currentGame.Alphabet = alphabet
		if Hangman.GetTryAttempt() == 0 {
			http.Redirect(w, r, "/FrLoose", http.StatusFound)
			return
		}
		if Hangman.CheckVictory() == true || currentGame.WordSelect == guessWord {
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
	Tg := template.Must(template.ParseFiles("./page/Fr/FrMedium.html"))
	if currentGame.WordSelect == "" {
		currentGame = Game(Hangman.GameInit("OriginalHangman/data/french/words2.txt"))
	}
	if r.Method == http.MethodPost {
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")
		Hangman.SelecCharac(guess)
		Hangman.SelecCharac(guessWord)
		currentGame.GuessedWord = Hangman.GetGuessedWord()
		currentGame.GuessedLetter = Hangman.GetGuessedLetter()
		currentGame.Try = Hangman.GetTryAttempt()
		currentGame.Alphabet = alphabet
		if Hangman.GetTryAttempt() == 0 {
			http.Redirect(w, r, "/FrLoose", http.StatusFound)
			return
		}
		if Hangman.CheckVictory() == true || currentGame.WordSelect == guessWord {
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
	Tg := template.Must(template.ParseFiles("./page/FR/FrHard.html"))
	if currentGame.WordSelect == "" {
		currentGame = Game(Hangman.GameInit("OriginalHangman/data/french/words3.txt"))
	}
	if r.Method == http.MethodPost {
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")
		Hangman.SelecCharac(guess)
		Hangman.SelecCharac(guessWord)
		currentGame.GuessedWord = Hangman.GetGuessedWord()
		currentGame.GuessedLetter = Hangman.GetGuessedLetter()
		currentGame.Try = Hangman.GetTryAttempt()
		currentGame.Alphabet = alphabet
		if Hangman.GetTryAttempt() == 0 {
			http.Redirect(w, r, "/FrLoose", http.StatusFound)
			return
		}
		if Hangman.CheckVictory() == true || currentGame.WordSelect == guessWord {
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
