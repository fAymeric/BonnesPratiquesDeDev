package internal

import (
	"html/template"
	"net/http"
)

func FrPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/Fr.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	currentGame = Game{}
	Tg.Execute(w, nil)
}
func FrWinpage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/FrWin.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Tg.Execute(w, nil)
}
func FrLoosepage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/FrLoose.html"))
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
