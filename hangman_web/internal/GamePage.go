package internal

import (
	"fmt"
	"hangman_web/internal/gameLogic"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
	"strings"
)

// GamePage handles the main game interface.
// It parses the language and difficulty from the URL, initializes the game if needed,
// processes guesses, and renders the game page.
func GamePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GamePage handler started")

	// Extract path and split to get language and difficulty
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	fmt.Println("URL Parts:", parts)

	// Validate the URL structure, must be /lang/difficulty
	if len(parts) != 2 {
		fmt.Println("Invalid URL format")
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	lang := parts[0]
	difficulty := parts[1]
	fmt.Println("Lang:", lang, " | Difficulty:", difficulty)

	// Validate supported languages and difficulty levels
	validLangs := map[string]bool{"Fr": true, "En": true, "Deutsch": true}
	validDiff := map[string]bool{"easy": true, "medium": true, "hard": true}
	if !validLangs[lang] || !validDiff[difficulty] {
		fmt.Println("Invalid language or difficulty")
		http.Error(w, "Langue ou difficulté invalide", http.StatusBadRequest)
		return
	}

	// Redirect to home if no user is logged in
	if shared.User == "" {
		fmt.Println("User not logged in")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Construct the path to the word file based on language and difficulty
	WordFile := fmt.Sprintf("internal/data/%s/words", lang)
	switch difficulty {
	case "easy":
		WordFile += ".txt"
	case "medium":
		WordFile += "2.txt"
	case "hard":
		WordFile += "3.txt"
	}
	fmt.Println("Word file:", WordFile)

	// If no word has been selected yet, initialize the game
	if shared.CurrentGame.WordSelect == "" {
		fmt.Println("Initialising game...")
		shared.CurrentGame = gameLogic.GameInit(WordFile)
	}

	// Handle guesses when the user submits a form
	if r.Method == http.MethodPost {
		fmt.Println("Handling POST")
		guess := r.FormValue("guess")
		guessWord := r.FormValue("guessWord")

		// Process both letter and full-word guesses
		gameLogic.SelecCharac(guess)
		gameLogic.SelecCharac(guessWord)

		// Update game state after guess
		shared.CurrentGame.GuessedWord = gameLogic.GetGuessedWord()
		shared.CurrentGame.GuessedLetter = gameLogic.GetGuessedLetter()
		shared.CurrentGame.Try = gameLogic.GetTryAttempt()
		shared.CurrentGame.Alphabet = shared.Alphabet

		// Check for loss condition
		if gameLogic.GetTryAttempt() == 0 {
			fmt.Println("User lost")
			RenderLoosePage(w, r, lang)
			return
		}

		// Check for win condition
		if gameLogic.CheckVictory() || shared.CurrentGame.WordSelect == guessWord {
			fmt.Println("User won")
			RenderWinPage(w, r, lang)
			shared.Score++
			gameLogic.UpdateScore(shared.User)
			return
		}
	}

	// Load localized texts based on the selected language
	localized := GetGameTexts(lang)

	// Prepare data for rendering the HTML template
	data := struct {
		shared.Game
		Texts map[string]string
	}{
		Game:  shared.CurrentGame,
		Texts: localized,
	}

	fmt.Println("Rendering template web/template/game.html")

	// Load and parse the game template
	tmpl, err := template.ParseFiles("web/template/game.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the game data
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Page rendered successfully")
}
