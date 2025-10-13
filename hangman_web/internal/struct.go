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
	Alphabet      string
}

var scores []ScoreEntry

var Score int
var user string
var currentGame Game
var alphabet = "abcdefghijklmnopqrstuvwxyz"

var word string
var letter string
var guessedWord []string
var guessedLetter []string
var tryAttempt int
var foundLetter bool
