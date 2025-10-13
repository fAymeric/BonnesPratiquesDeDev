package internal

import (
	"html/template"
	"net/http"
)

func EnPage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/En/En.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	currentGame = Game{}
	Tg.Execute(w, nil)
}
func EnWinpage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/En/EnWin.html"))
	if user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Tg.Execute(w, nil)
}
func EnLoosepage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/En/EnLoose.html"))
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
