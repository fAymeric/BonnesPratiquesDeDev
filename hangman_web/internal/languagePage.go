package internal

import (
	"fmt"
	"hangman_web/internal/shared"
	"html/template"
	"net/http"
)

func LanguagePage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Language.html"))
	if r.Method != "POST" && shared.User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	newUser := r.FormValue("name")
	if shared.User == "" {
		shared.User = r.FormValue("name")
		fmt.Println("user : ", shared.User)
		err := Tg.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	if shared.User != "" && newUser != shared.User {
		shared.User = newUser
		shared.Score = 0
		fmt.Println("user : ", shared.User)
		err := Tg.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
