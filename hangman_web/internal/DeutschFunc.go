package internal

import (
	"html/template"
	"net/http"
)

func DeutschPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Deutsch.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	currentGame = Game{}
	Tg.Execute(w, nil)
}
func DeutschWinpage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Deutsch/DeutschWin.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Tg.Execute(w, nil)
}
func DeutschLoosepage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Deutsch/DeutschLoose.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	looseData := struct {
		WordSelect string
	}{
		WordSelect: currentGame.WordSelect,
	}
	Tg.Execute(w, looseData)
}
