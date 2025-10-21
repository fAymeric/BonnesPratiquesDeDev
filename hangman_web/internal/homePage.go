package internal

import (
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	// Only append the score if it's not already in the scores list
	for _, scoreEntry := range shared.Scores {
		if scoreEntry.Pseudo == shared.User {
			break
		}
	}

	// Render the homepage template
	Th := template.Must(template.ParseFiles("web/template/index.html"))
	err := Th.Execute(w, struct {
		Scores []shared.ScoreEntry
	}{
		Scores: shared.Scores,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
