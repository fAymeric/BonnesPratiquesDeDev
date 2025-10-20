package internal

import (
	"fmt"
	"hangman_web/internal/gameLogic"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
	"strings"
)

func GamePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("➡️ GamePage handler started")

	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	fmt.Println("➡️ URL Parts:", parts)

	if len(parts) != 2 {
		fmt.Println("❌ Invalid URL format")
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	lang := parts[0]
	difficulty := parts[1]

	fmt.Println("➡️ Lang:", lang, " | Difficulty:", difficulty)

	validLangs := map[string]bool{"Fr": true, "En": true, "Deutsch": true}
	validDiff := map[string]bool{"easy": true, "medium": true, "hard": true}

	if !validLangs[lang] || !validDiff[difficulty] {
		fmt.Println("❌ Invalid language or difficulty")
		http.Error(w, "Langue ou difficulté invalide", http.StatusBadRequest)
		return
	}

	if shared.User == "" {
		fmt.Println("❌ User not logged in")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Construire le chemin vers les mots
	WordFile := fmt.Sprintf("internal/data/%s/words", lang)
	switch difficulty {
	case "easy":
		WordFile += ".txt"
	case "medium":
		WordFile += "2.txt"
	case "hard":
		WordFile += "3.txt"
	}

	fmt.Println("➡️ Word file:", WordFile)

	if shared.CurrentGame.WordSelect == "" {
		fmt.Println("➡️ Initialising game...")
		shared.CurrentGame = gameLogic.GameInit(WordFile)
	}

	if r.Method == http.MethodPost {
		fmt.Println("➡️ Handling POST")
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")
		gameLogic.SelecCharac(guess)
		gameLogic.SelecCharac(guessWord)

		shared.CurrentGame.GuessedWord = gameLogic.GetGuessedWord()
		shared.CurrentGame.GuessedLetter = gameLogic.GetGuessedLetter()
		shared.CurrentGame.Try = gameLogic.GetTryAttempt()
		shared.CurrentGame.Alphabet = shared.Alphabet

		if gameLogic.GetTryAttempt() == 0 {
			fmt.Println("➡️ User lost")
			RenderLoosePage(w, r, lang)
			return
		}
		if gameLogic.CheckVictory() || shared.CurrentGame.WordSelect == guessWord {
			fmt.Println("🏆 User won")
			RenderWinPage(w, r, lang)
			shared.Score++
			gameLogic.UpdateScore(shared.User)
			return
		}
	}

	// Données localisées
	localized := GetGameTexts(lang)

	// Données pour le template
	data := struct {
		shared.Game
		Texts map[string]string
	}{
		Game:  shared.CurrentGame,
		Texts: localized,
	}

	fmt.Println("➡️ Rendering template web/template/game.html")

	tmpl, err := template.ParseFiles("web/template/game.html")
	if err != nil {
		fmt.Println("❌ Error parsing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("❌ Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("✅ Page rendered successfully")
}
