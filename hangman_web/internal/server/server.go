package server

import (
	"hangman_web/internal"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", internal.Homepage)
	http.HandleFunc("/EnWin", internal.EnWinpage)
	http.HandleFunc("/EnLoose", internal.EnLoosepage)
	http.HandleFunc("/DeutschWin", internal.DeutschWinpage)
	http.HandleFunc("/DeutschLoose", internal.DeutschLoosepage)
	http.HandleFunc("/FrWin", internal.FrWinpage)
	http.HandleFunc("/FrLoose", internal.FrLoosepage)
	http.HandleFunc("/language", internal.LanguagePage)
	http.HandleFunc("/Fr", internal.FrPage)
	http.HandleFunc("/En", internal.EnPage)
	http.HandleFunc("/Deutsch", internal.DeutschPage)
	http.HandleFunc("/Fr/easy", internal.GamePage)
	http.HandleFunc("/Fr/medium", internal.GamePage)
	http.HandleFunc("/Fr/hard", internal.GamePage)
	http.HandleFunc("/En/easy", internal.GamePage)
	http.HandleFunc("/En/medium", internal.GamePage)
	http.HandleFunc("/En/hard", internal.GamePage)
	http.HandleFunc("/Deutsh/easy", internal.GamePage)
	http.HandleFunc("/Deutsh/medium", internal.GamePage)
	http.HandleFunc("/Deutsh/hard", internal.GamePage)
	http.HandleFunc("/JS/", internal.FilesHandler)
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))
	http.ListenAndServe(":8080", nil)
}
