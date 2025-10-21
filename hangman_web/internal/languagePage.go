package internal

import (
	"fmt"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
)

// LanguagePage handles the language selection page.
// It also initializes user data if a new user submits their name.
func LanguagePage(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template for the language selection page
	Tg := template.Must(template.ParseFiles("web/template/language.html"))

	// Redirect to home if user is not logged in
	if r.Method != http.MethodPost && shared.User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse form values from POST request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := r.FormValue("name")

	// If a new username is submitted and different from the current user
	if newUser != "" && newUser != shared.User {
		// Update user info
		shared.User = newUser
		shared.Score = 0

		// Reset game-related data
		shared.CurrentGame = shared.Game{}
		shared.Word = ""
		shared.GuessedLetter = []string{}
		shared.GuessedWord = []string{}
		shared.TryAttempt = 0

		fmt.Println("New user:", shared.User)
	} else {
		// If no new user, still reset the game state
		shared.CurrentGame = shared.Game{}
		shared.Word = ""
		shared.GuessedLetter = []string{}
		shared.GuessedWord = []string{}
		shared.TryAttempt = 0
	}

	// Render the language selection template
	err = Tg.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
