package internal

import (
	"html/template"
	"net/http"
)

func DeutschPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Deutsch/Deutsch.html"))
	if User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	CurrentGame = Game{}
	err := Tg.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func DeutschWinpage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Win.html"))
	if User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := Tg.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func DeutschLoosepage(w http.ResponseWriter, r *http.Request) {
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
	err := Tg.Execute(w, looseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
