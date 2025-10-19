package internal

import (
	"html/template"
	"net/http"
)

func FrPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Fr/Fr.html"))
	if User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	CurrentGame = Game{}
	Tg.Execute(w, nil)
}
func FrWinpage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Win.html"))
	if User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Tg.Execute(w, nil)
}
func FrLoosepage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Loose.html"))
	if User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	looseData := struct {
		WordSelect string
	}{
		WordSelect: CurrentGame.WordSelect,
	}
	Tg.Execute(w, looseData)
}
