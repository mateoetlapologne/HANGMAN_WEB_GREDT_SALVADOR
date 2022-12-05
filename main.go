package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username   string
	Difficulty string
	TryNumber  int
	LetterTry  []string
	Success    bool
}

var details = User{
	TryNumber: 10,
	LetterTry: []string{},
	Success:   false,
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
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	details.Success = true
	tmpl1.Execute(w, details)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	details.Success = false
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	// details.TryNumber -= 1
	// details.LetterTry = append(details.LetterTry, r.FormValue("LetterTry"))
	fmt.Println(details.Username)
	fmt.Println(details.Difficulty)
	fmt.Println(details.TryNumber)
	fmt.Println(details.LetterTry)
	tmpl1 := template.Must(template.ParseFiles("game.html"))
	tmpl1.Execute(w, details)
}
