package internal

import (
	"html/template"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	// Only append the score if it's not already in the scores list
	for _, scoreEntry := range scores {
		if scoreEntry.Pseudo == user {
			break
		}
	}

	// Render the homepage template
	Th := template.Must(template.ParseFiles("web/template/SiteH.html"))
	Th.Execute(w, struct {
		Scores []ScoreEntry
	}{
		Scores: scores,
	})
}
