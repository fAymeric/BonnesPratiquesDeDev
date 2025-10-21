package internal

import (
	"bytes"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
)

// RenderLoosePage renders the "lose" page when the player runs out of tries.
// It uses localized text depending on the selected language and displays the correct word.
func RenderLoosePage(w http.ResponseWriter, r *http.Request, language string) {
	// Load the HTML template for the lose page.
	tmpl := template.Must(template.ParseFiles("web/template/loose.html"))

	var data shared.WinLoosePageData

	// Select localized text content based on the chosen language
	switch language {
	case "Fr":
		data = shared.WinLoosePageData{
			Language:   "Fr",
			PageTitle:  "Page de défaite",
			TextLoose:  "Tu es mort... Au moins ta chair aura servi à quelque chose. Le mot était",
			LinkHome:   "Retour au menu principal",
			LinkReplay: "Envie de rejouer ?",
		}
	case "En":
		data = shared.WinLoosePageData{
			Language:   "En",
			PageTitle:  "Lose page",
			TextLoose:  "You Died... At least your flesh has been useful. The word was",
			LinkHome:   "Back to main menu",
			LinkReplay: "Want to play again?",
		}
	case "Deutsch":
		data = shared.WinLoosePageData{
			Language:   "Deutsch",
			PageTitle:  "Lose Seite",
			TextLoose:  "Du bist gestorben ... Zumindest war dein Fleisch nützlich. Das Wort war",
			LinkHome:   "Zurück zum Hauptmenü",
			LinkReplay: "Willst du noch einmal spielen?",
		}
	default:
		// Return error if language is unsupported
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	// Add the correct word to the data to display to the user
	data.WordSelect = shared.CurrentGame.WordSelect

	// Render the template into a buffer first
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, "Internal error while rendering page", http.StatusInternalServerError)
		return
	}

	// If successful, write the buffer contents to the response
	buf.WriteTo(w)
}
