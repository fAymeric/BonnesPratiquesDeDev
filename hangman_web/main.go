package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)
import "html/template"

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
	startServer()
}

func startServer() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/EnWin", EnWinpage)
	http.HandleFunc("/EnLoose", EnLoosepage)
	http.HandleFunc("/DeutschWin", DeutschWinpage)
	http.HandleFunc("/DeutschLoose", DeutschLoosepage)
	http.HandleFunc("/FrWin", FrWinpage)
	http.HandleFunc("/FrLoose", FrLoosepage)
	http.HandleFunc("/language", LanguagePage)
	http.HandleFunc("/Fr", FrPage)
	http.HandleFunc("/En", EnPage)
	http.HandleFunc("/Deutsch", DeutschPage)
	http.HandleFunc("/Fr/easy", FrEasyPage)
	http.HandleFunc("/Fr/medium", FrMediumPage)
	http.HandleFunc("/Fr/hard", FrHardPage)
	http.HandleFunc("/En/easy", EnEasyPage)
	http.HandleFunc("/En/medium", EnMediumPage)
	http.HandleFunc("/En/hard", EnHardPage)
	http.HandleFunc("/Deutsch/easy", DeutschEasyPage)
	http.HandleFunc("/Deutsch/medium", DeutschMediumPage)
	http.HandleFunc("/Deutsch/hard", DeutschHardPage)
	http.HandleFunc("/JS/", filesHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)

}

func Homepage(w http.ResponseWriter, r *http.Request) {
	// Only append the score if it's not already in the scores list
	for _, scoreEntry := range scores {
		if scoreEntry.Pseudo == user {
			break
		}
	}

	// Render the homepage template
	Th := template.Must(template.ParseFiles("./page/SiteH.html"))
	Th.Execute(w, struct {
		Scores []ScoreEntry
	}{
		Scores: scores,
	})
}

func LanguagePage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Language.html"))
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
