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
}

var details = User{
	Username:   "",
	Difficulty: "",
	Success:    false,
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")

	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	h := hangman.HangManData{}
	h.Init()
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
	h.Init()
	h.Game(r.FormValue("LetterTry"))
	fmt.Println("Lettre entrée ", r.FormValue("LetterTry"))
	fmt.Println("Mot à trouver ", h.Word)
	fmt.Println("Affichage ", h.ToFind)
	fmt.Println("Lettres connues ", h.KnownLetters)
	fmt.Println("Il vous reste ", h.Attempts, " tentatives")
	fmt.Println("Vous avez déjà tenté ")
	for _, v := range h.TriedLetters {
		fmt.Println(v)
	}
	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
