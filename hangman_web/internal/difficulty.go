package internal

import (
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
	"strings"
)

// DifficultyPage handles rendering the difficulty selection page.
// If the user is not logged in, they are redirected to the home page.
// It also loads localized text content based on the selected language.
func DifficultyPage(w http.ResponseWriter, r *http.Request) {
	// Extract language code from URL ("/En" = "En")
	lang := strings.TrimPrefix(r.URL.Path, "/")

	// Ensure a user is logged in before accessing this page
	if shared.User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Localized text data for supported languages
	localizedTexts := map[string]map[string]string{
		"Fr": {
			"title":  "Difficulté",
			"easy":   "Facile",
			"medium": "Moyen",
			"hard":   "Difficile",
			"back":   "Retour",
		},
		"En": {
			"title":  "Difficulty",
			"easy":   "Easy",
			"medium": "Medium",
			"hard":   "Hard",
			"back":   "Back",
		},
		"Deutsch": {
			"title":  "Schwierigkeit",
			"easy":   "Einfach",
			"medium": "Durchschnitt",
			"hard":   "Schwierig",
			"back":   "Zurück",
		},
	}

	// Validate the requested language and return error if unsupported
	texts, ok := localizedTexts[lang]
	if !ok {
		http.Error(w, "Invalid language", http.StatusBadRequest)
		return
	}

	// Prepare data to pass into the HTML template
	data := shared.DifficultyPageData{
		Lang:  lang,
		Texts: texts,
	}

	// Parse the HTML template for the difficulty page
	tmpl, err := template.ParseFiles("web/template/difficulty.html")
	if err != nil {
		http.Error(w, "Template loading error", http.StatusInternalServerError)
		return
	}

	// Render the template with data and handle any execution errors
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}
