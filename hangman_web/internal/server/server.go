package server

import "net/http"

func StartServer() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/EnWin", EnWinpage)
	http.HandleFunc("/EnLoose", EnLoosepage)
	http.HandleFunc("/DeutschWin", DeutschWinpage)
	http.HandleFunc("/DeutschLoose", DeutschLoosepage)
	http.HandleFunc("/FrWin", FrWinpage)
	http.HandleFunc("/FrLoose", FrLoosepage)
	http.HandleFunc("/language", LanguagePage)
	http.HandleFunc("/Fr", FrPage)
	http.HandleFunc("/En", EnPage)
	http.HandleFunc("/Deutsch", DeutschPage)
	http.HandleFunc("/Fr/easy", FrEasyPage)
	http.HandleFunc("/Fr/medium", FrMediumPage)
	http.HandleFunc("/Fr/hard", FrHardPage)
	http.HandleFunc("/En/easy", EnEasyPage)
	http.HandleFunc("/En/medium", EnMediumPage)
	http.HandleFunc("/En/hard", EnHardPage)
	http.HandleFunc("/Deutsch/easy", DeutschEasyPage)
	http.HandleFunc("/Deutsch/medium", DeutschMediumPage)
	http.HandleFunc("/Deutsch/hard", DeutschHardPage)
	http.HandleFunc("/JS/", filesHandler)
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))
	http.ListenAndServe(":8080", nil)

}
