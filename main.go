package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Success  bool
}

var details = User{
	Username: "none",
	Success:  false,
}

func main() {
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	//gestion html
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.ListenAndServe(":80", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	details.Username = r.FormValue("difficulte")
	fmt.Println(details.Username)
	details.Success = true
	tmpl1.Execute(w, details)

}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("game.html"))
	tmpl1.Execute(w, details)
}
