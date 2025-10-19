package internal

func GetGameTexts(lang string) map[string]string {
	switch lang {
	case "Fr":
		return map[string]string{
			"title":       "Le Pendu",
			"subtitle":    "Et maintenant jouons...",
			"wordLabel":   "Mot",
			"usedLetters": "Lettres déjà trouvées",
			"back":        "Retour",
		}
	case "En":
		return map[string]string{
			"title":       "The Hangman",
			"subtitle":    "It's time to play...",
			"wordLabel":   "Word",
			"usedLetters": "Already used Letters",
			"back":        "Back",
		}
	case "Deutsch":
		return map[string]string{
			"title":       "Der Gehängte",
			"subtitle":    "Und jetzt lasst uns spielen...",
			"wordLabel":   "Wort",
			"usedLetters": "Brief bereits verschickt",
			"back":        "zurück",
		}
	default:
		return map[string]string{}
	}
}
