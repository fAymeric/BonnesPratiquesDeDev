package server

import (
	"hangman_web/internal"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", internal.Homepage)
	http.HandleFunc("/language", internal.LanguagePage)
	http.HandleFunc("/Fr", internal.DifficultyPage)
	http.HandleFunc("/En", internal.DifficultyPage)
	http.HandleFunc("/Deutsch", internal.DifficultyPage)
	http.HandleFunc("/Fr/easy", internal.GamePage)
	http.HandleFunc("/Fr/medium", internal.GamePage)
	http.HandleFunc("/Fr/hard", internal.GamePage)
	http.HandleFunc("/En/easy", internal.GamePage)
	http.HandleFunc("/En/medium", internal.GamePage)
	http.HandleFunc("/En/hard", internal.GamePage)
	http.HandleFunc("/Deutsch/easy", internal.GamePage)
	http.HandleFunc("/Deutsch/medium", internal.GamePage)
	http.HandleFunc("/Deutsch/hard", internal.GamePage)
	http.HandleFunc("/JS/", internal.FilesHandler)
	http.HandleFunc("/FrWin", func(w http.ResponseWriter, r *http.Request) {
		internal.RenderWinPage(w, r, "Fr")
	})
	http.HandleFunc("/FrLoose", func(w http.ResponseWriter, r *http.Request) {
		internal.RenderLoosePage(w, r, "Fr")
	})

	http.HandleFunc("/EnWin", func(w http.ResponseWriter, r *http.Request) {
		internal.RenderWinPage(w, r, "En")
	})
	http.HandleFunc("/EnLoose", func(w http.ResponseWriter, r *http.Request) {
		internal.RenderLoosePage(w, r, "En")
	})

	http.HandleFunc("/DeutschWin", func(w http.ResponseWriter, r *http.Request) {
		internal.RenderWinPage(w, r, "Deutsch")
	})
	http.HandleFunc("/DeutschLoose", func(w http.ResponseWriter, r *http.Request) {
		internal.RenderLoosePage(w, r, "Deutsch")
	})
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))
	http.ListenAndServe(":8080", nil)
}
