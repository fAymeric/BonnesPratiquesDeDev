package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type ScoreEntry struct {
	Pseudo     string
	Scoreperso int
}

var scores []ScoreEntry

type Game struct {
	WordSelect    string
	AppearLetter  []string
	GuessedLetter []string
	Try           int
	GuessedWord   []string
	Alphabet      string
}

var Score int
var user string
var currentGame Game
var alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
	server.startServer()
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	// Only append the score if it's not already in the scores list
	for _, scoreEntry := range scores {
		if scoreEntry.Pseudo == user {
			break
		}
	}

	// Render the homepage template
	Th := template.Must(template.ParseFiles("../web/template/SiteH.html"))
	Th.Execute(w, struct {
		Scores []ScoreEntry
	}{
		Scores: scores,
	})
}

func LanguagePage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("../web/template/Language.html"))
	if r.Method != "POST" && user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	newUser := r.FormValue("name")
	if user == "" {
		user = r.FormValue("name")
		fmt.Println("user : ", user)
		Tg.Execute(w, nil)
	}
	if user != "" && newUser != user {
		user = newUser
		Score = 0
		fmt.Println("user : ", user)
		Tg.Execute(w, nil)
	}
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path
	filePath := "." + filename

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "File not found", 404)
		return
	}

	var fileType string
	switch filepath.Ext(filename) {
	case ".js":
		fileType = "application/javascript"
	}

	w.Header().Set("Content-Type", fileType)
	w.Write(file)
}
