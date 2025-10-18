package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

func LanguagePage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Language.html"))
	if r.Method != "POST" && user == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	newUser := r.FormValue("name")
	if user == "" {
		user = r.FormValue("name")
		fmt.Println("user : ", user)
		err := Tg.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	if user != "" && newUser != user {
		user = newUser
		Score = 0
		fmt.Println("user : ", user)
		err := Tg.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
