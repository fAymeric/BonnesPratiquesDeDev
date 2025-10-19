package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

func LanguagePage(w http.ResponseWriter, r *http.Request) {
	Tg := template.Must(template.ParseFiles("web/template/Language.html"))
	if r.Method != "POST" && User == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	newUser := r.FormValue("name")
	if User == "" {
		User = r.FormValue("name")
		fmt.Println("user : ", User)
		err := Tg.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	if User != "" && newUser != User {
		User = newUser
		Score = 0
		fmt.Println("user : ", User)
		err := Tg.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
