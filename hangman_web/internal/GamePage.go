package internal

import (
	"fmt"
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

	if User == "" {
		fmt.Println("❌ User not logged in")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Construire le chemin vers les mots
	wordFile := fmt.Sprintf("internal/data/%s/words", lang)
	switch difficulty {
	case "easy":
		wordFile += ".txt"
	case "medium":
		wordFile += "2.txt"
	case "hard":
		wordFile += "3.txt"
	}

	fmt.Println("➡️ Word file:", wordFile)

	if CurrentGame.WordSelect == "" {
		fmt.Println("➡️ Initialising game...")
		CurrentGame = GameInit(wordFile)
	}

	if r.Method == http.MethodPost {
		fmt.Println("➡️ Handling POST")
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")
		SelecCharac(guess)
		SelecCharac(guessWord)

		CurrentGame.GuessedWord = GetGuessedWord()
		CurrentGame.GuessedLetter = GetGuessedLetter()
		CurrentGame.Try = GetTryAttempt()
		CurrentGame.Alphabet = Alphabet

		if GetTryAttempt() == 0 {
			fmt.Println("➡️ User lost")
			RenderLoosePage(w, r, lang)
			return
		}
		if CheckVictory() || CurrentGame.WordSelect == guessWord {
			fmt.Println("🏆 User won")
			RenderWinPage(w, r, lang)
			Score++
			UpdateScore(User)
			return
		}
	}

	// Données localisées
	localized := GetGameTexts(lang)

	// Données pour le template
	data := struct {
		Game
		Texts map[string]string
	}{
		Game:  CurrentGame,
		Texts: localized,
	}

	fmt.Println("➡️ Rendering template web/template/GamePage.html")

	tmpl, err := template.ParseFiles("web/template/Game.html")
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
