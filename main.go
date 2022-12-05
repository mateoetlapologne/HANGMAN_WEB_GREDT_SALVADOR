// main.go

package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username   string
	Difficulty string
}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := User{
			Username:   r.FormValue("username"),
			Difficulty: r.FormValue("difficulty"),
		}
		tmpl.Execute(w, details)
	})

	tmp2 := template.Must(template.ParseFiles("game.html"))
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := User{
			Username:   r.FormValue("username"),
			Difficulty: r.FormValue("difficulty"),
		}
		tmp2.Execute(w, details)
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
