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
var user string
var currentGame Game
var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var word string
var guessedWord []string
var guessedLetter []string
var tryAttempt int
var foundLetter bool
