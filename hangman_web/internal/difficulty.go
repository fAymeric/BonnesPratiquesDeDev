package internal

import (
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
	"strings"
)

func DifficultyPage(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, "/")

	if shared.User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

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

	texts, ok := localizedTexts[lang]
	if !ok {
		http.Error(w, "Langue invalide", http.StatusBadRequest)
		return
	}

	data := shared.DifficultyPageData{
		Lang:  lang,
		Texts: texts,
	}

	tmpl, err := template.ParseFiles("web/template/Difficulty.html")
	if err != nil {
		http.Error(w, "Erreur chargement template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur exécution template", http.StatusInternalServerError)
	}
}
