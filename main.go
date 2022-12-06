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
	Username:   "",
	Difficulty: "",
	TryNumber:  11,
	LetterTry:  []string{},
	Success:    false,
}

func main() {
	//démarage serveur
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
	tmpl1 := template.Must(template.ParseFiles("templates/index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	//details.Success = true
	tmpl1.Execute(w, details)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	details.TryNumber--
	details.LetterTry = append(details.LetterTry, r.FormValue("LetterTry"))

	//debug pour voir les valeurs
	fmt.Println(details.Username)
	fmt.Println(details.Difficulty)
	fmt.Println(details.TryNumber)
	fmt.Println(details.LetterTry)

	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
