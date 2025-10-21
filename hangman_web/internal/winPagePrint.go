package internal

import (
	"bytes"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
)

// RenderWinPage renders the victory page when the player wins the game.
// It uses localized content based on the selected language.
func RenderWinPage(w http.ResponseWriter, r *http.Request, language string) {
	// Load the HTML template for the win page.
	tmpl := template.Must(template.ParseFiles("web/template/win.html"))

	var data shared.WinLoosePageData

	// Select localized strings based on the language parameter
	switch language {
	case "Fr":
		data = shared.WinLoosePageData{
			Language:   "Fr",
			PageTitle:  "Page de victoire",
			TextWin:    "Bien joué, tu as réussi à t'échapper... Mais pour combien de temps...",
			LinkHome:   "Retour au menu principal",
			LinkReplay: "Envie de rejouer ?",
		}
	case "En":
		data = shared.WinLoosePageData{
			Language:   "En",
			PageTitle:  "Victory Page",
			TextWin:    "Well done, you managed to escape... But for how long...",
			LinkHome:   "Back to main menu",
			LinkReplay: "Want to play again?",
		}
	case "Deutsch":
		data = shared.WinLoosePageData{
			Language:   "Deutsch",
			PageTitle:  "Siegesseite",
			TextWin:    "Gut gemacht, du bist entkommen... Aber wie lange noch...",
			LinkHome:   "Zurück zum Hauptmenü",
			LinkReplay: "Noch eine Runde?",
		}
	default:
		// Return an error if the language is not supported
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	// Render the template into a buffer
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, "Internal error rendering the page", http.StatusInternalServerError)
		return
	}

	// If no error, write the fully rendered page to the response
	buf.WriteTo(w)
}
