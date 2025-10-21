package internal

// GetGameTexts returns localized strings for the main game interface,
// based on the selected language code
func GetGameTexts(lang string) map[string]string {
	switch lang {
	case "Fr":
		// French translations
		return map[string]string{
			"title":       "Le Pendu",
			"subtitle":    "Et maintenant jouons...",
			"wordLabel":   "Mot",
			"usedLetters": "Lettres déjà trouvées",
			"back":        "Retour",
		}
	case "En":
		// English translations
		return map[string]string{
			"title":       "The Hangman",
			"subtitle":    "It's time to play...",
			"wordLabel":   "Word",
			"usedLetters": "Already used Letters",
			"back":        "Back",
		}
	case "Deutsch":
		// German translations
		return map[string]string{
			"title":       "Der Gehängte",
			"subtitle":    "Und jetzt lasst uns spielen...",
			"wordLabel":   "Wort",
			"usedLetters": "Brief bereits verschickt",
			"back":        "zurück",
		}
	default:
		// Return an empty map if the language is unsupported
		return map[string]string{}
	}
}
