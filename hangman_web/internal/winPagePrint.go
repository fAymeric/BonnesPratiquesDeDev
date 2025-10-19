package internal

import (
	"bytes"
	"html/template"
	"net/http"
)

func RenderWinPage(w http.ResponseWriter, r *http.Request, language string) {
	tmpl := template.Must(template.ParseFiles("web/template/Win.html"))

	var data WinLoosePageData

	switch language {
	case "fr":
		data = WinLoosePageData{
			Language:   "fr",
			PageTitle:  "Page de victoire",
			TextWin:    "Bien joué, tu as réussi à t'échapper... Mais pour combien de temps...",
			LinkHome:   "Retour au menu principal",
			LinkReplay: "Envie de rejouer ?",
		}
	case "en":
		data = WinLoosePageData{
			Language:   "en",
			PageTitle:  "Victory Page",
			TextWin:    "Well done, you managed to escape... But for how long...",
			LinkHome:   "Back to main menu",
			LinkReplay: "Want to play again?",
		}
	case "de":
		data = WinLoosePageData{
			Language:   "de",
			PageTitle:  "Siegesseite",
			TextWin:    "Gut gemacht, du bist entkommen... Aber wie lange noch...",
			LinkHome:   "Zurück zum Hauptmenü",
			LinkReplay: "Noch eine Runde?",
		}
	default:
		http.Error(w, "Langue non prise en charge", http.StatusBadRequest)
		return
	}

	// Rendu dans un buffer d'abord
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, "Erreur interne lors du rendu de la page", http.StatusInternalServerError)
		return
	}

	// Écriture dans la réponse seulement si tout s’est bien passé
	buf.WriteTo(w)
}
