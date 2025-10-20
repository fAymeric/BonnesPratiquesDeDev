package shared

type ScoreEntry struct {
	Pseudo     string
	Scoreperso int
}

type Game struct {
	WordSelect    string
	AppearLetter  []string
	GuessedLetter []string
	Try           int
	GuessedWord   []string
	Alphabet      []string
}

// Global Data
var (
	Scores      []ScoreEntry
	Score       int
	User        string
	CurrentGame Game
	Alphabet    = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	Word          string
	GuessedWord   []string
	GuessedLetter []string
	TryAttempt    int
	FoundLetter   bool
)

// Structure pour les pages de victoire/défaite
type WinLoosePageData struct {
	Language   string
	PageTitle  string
	TextWin    string
	TextLoose  string
	LinkHome   string
	LinkReplay string
	WordSelect string
}

// Structure pour la page de sélection de difficulté (multi-langue)
type DifficultyPageData struct {
	Lang  string
	Texts map[string]string
}
