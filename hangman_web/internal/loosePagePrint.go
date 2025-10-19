package internal

import (
	"bytes"
	"html/template"
	"net/http"
)

func RenderLoosePage(w http.ResponseWriter, r *http.Request, language string) {
	tmpl := template.Must(template.ParseFiles("web/template/Loose.html"))

	var data WinLoosePageData

	// Texte selon la langue
	switch language {
	case "Fr":
		data = WinLoosePageData{
			Language:   "Fr",
			PageTitle:  "Page de défaite",
			TextLoose:  "Tu es mort... Au moins ta chair aura servi à quelque chose. Le mot était",
			LinkHome:   "Retour au menu principal",
			LinkReplay: "Envie de rejouer ?",
		}
	case "En":
		data = WinLoosePageData{
			Language:   "En",
			PageTitle:  "Lose page",
			TextLoose:  "You Died... At least your flesh has been useful. The word was",
			LinkHome:   "Back to main menu",
			LinkReplay: "Want to play again?",
		}
	case "Deutsch":
		data = WinLoosePageData{
			Language:   "Deutsch",
			PageTitle:  "Lose Seite",
			TextLoose:  "Du bist gestorben ... Zumindest war dein Fleisch nützlich. Das Wort war",
			LinkHome:   "Zurück zum Hauptmenü",
			LinkReplay: "Willst du noch einmal spielen?",
		}
	default:
		http.Error(w, "Langue non prise en charge", http.StatusBadRequest)
		return
	}

	data.WordSelect = CurrentGame.WordSelect

	// Rendu dans un buffer d'abord
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, "Erreur interne lors du rendu de la page", http.StatusInternalServerError)
		return
	}

	// Si tout est OK, écriture
	buf.WriteTo(w)
}
