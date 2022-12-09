package main

import (
	"fmt"
	"html/template"
	"net/http"
	hangman "serv-hangman/packages"
)

type User struct {
	Username    string
	Difficulty  string
	Word        string
	ToFind      string
	Attempts    int
	LetterTry   []string
	LetterKnown []string
}

var h hangman.HangManData

var details = User{
	Username:   "",
	Difficulty: "",
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	h.Init()
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
	h.Game(r.FormValue("LetterTry"))
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	details.Word = h.Word
	details.ToFind = h.ToFind
	details.Attempts = h.Attempts
	details.LetterTry = h.TriedLetters
	details.LetterKnown = h.KnownLetters
	fmt.Println("Lettre entrée ", r.FormValue("LetterTry"))
	fmt.Println("Mot à trouver ", h.ToFind)
	fmt.Println("Affichage ", h.Word)
	fmt.Println("Lettres connues ", h.KnownLetters)
	fmt.Println("Vous avez déjà tenté ", h.TriedLetters)
	fmt.Println("Il vous reste ", h.Attempts, " tentatives")
	for _, v := range h.TriedLetters {
		fmt.Println(v)
	}
	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
