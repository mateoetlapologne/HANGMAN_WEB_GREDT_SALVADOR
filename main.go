// main.go

package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username   string
	Difficulty string
}

func main() {
	tmpl := template.Must(template.ParseGlob("*.html"))
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
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
	http.ListenAndServe(":80", nil)
}
