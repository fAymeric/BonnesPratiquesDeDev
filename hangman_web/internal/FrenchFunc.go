package internal

import (
	"html/template"
	"net/http"
)

func FrPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Fr.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	currentGame = Game{}
	Tg.Execute(w, nil)
}
func FrWinpage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Fr/FrWin.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Tg.Execute(w, nil)
}
func FrLoosepage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("./page/Fr/FrLoose.html"))
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
