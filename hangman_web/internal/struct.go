package internal

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

var scores []ScoreEntry

var Score int
var User string
var CurrentGame Game
var Alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var Word string
var GuessedWord []string
var GuessedLetter []string
var TryAttempt int
var FoundLetter bool

type WinLoosePageData struct {
	Language   string
	PageTitle  string
	TextWin    string
	TextLoose  string
	LinkHome   string
	LinkReplay string
	WordSelect string
}
