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
	WordDisplay string
	TryNumber   int
	LetterTry   []string
	Success     bool
	Data        hangman.HangManData
}

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
	h := hangman.HangManData{}
	details.Data = h.Init()
	// details.Word = h.ToFind
	// details.WordDisplay = h.Word

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
	h := hangman.HangManData{}
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	h.Game(r.FormValue("LetterTry"), details.Data)
	fmt.Println("Lettre entrée ", r.FormValue("LetterTry"))
	fmt.Println("Mot à trouver ", details.Data.ToFind)
	fmt.Println("Affichage ", details.Data.Word)
	fmt.Println("Lettres connues ", details.Data.KnownLetters)
	fmt.Println("Vous avez déjà tenté ", details.Data.TriedLetters)
	fmt.Println("Il vous reste ", details.Data.Attempts, " tentatives")
	for _, v := range h.TriedLetters {
		fmt.Println(v)
	}
	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
