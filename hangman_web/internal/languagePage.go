package internal

import (
	"fmt"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
)

func LanguagePage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/language.html"))

	// Redirection vers "/" si l'utilisateur n'est pas encore défini et la méthode n'est pas POST
	if r.Method != http.MethodPost && shared.User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := r.FormValue("name")

	// Si un nouveau nom est soumis
	if newUser != "" && newUser != shared.User {
		shared.User = newUser
		shared.Score = 0

		// Réinitialisation du jeu
		shared.CurrentGame = shared.Game{}
		shared.Word = ""
		shared.GuessedLetter = []string{}
		shared.GuessedWord = []string{}
		shared.TryAttempt = 0

		fmt.Println("Nouvel utilisateur :", shared.User)
	} else {
		shared.CurrentGame = shared.Game{}
		shared.Word = ""
		shared.GuessedLetter = []string{}
		shared.GuessedWord = []string{}
		shared.TryAttempt = 0
	}

	err = Tg.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
