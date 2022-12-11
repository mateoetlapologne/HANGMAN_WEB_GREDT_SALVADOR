package main

import (
	"fmt"
	"html/template"
	"net/http"
	hangman "serv-hangman/packages"
	"strings"
	"unicode"
)

type User struct {
	Username    string
	Difficulty  string
	Word        string
	ToFind      string
	Win         int
	Lose        int
	Attempts    int
	LetterTry   []string
	LetterKnown []string
	Display     string
}

var h hangman.HangManData

var details = User{
	Username:   "",
	Difficulty: "",
	Attempts:   11,
	Win:        5,
	Lose:       3,
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

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("templates/index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
		return
	}
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	tmpl1.Execute(w, details)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	details.Username = r.FormValue("Username")
	details.Difficulty = r.FormValue("Difficulty")
	if details.Attempts == 11 {
		h = hangman.HangManData{}
		if details.Difficulty == "Facile" {
			h.Init("words.txt")
		} else if details.Difficulty == "Moyenne" {
			h.Init("words2.txt")
		} else if details.Difficulty == "Difficile" {
			h.Init("words3.txt")
		}
	}
	if IsLetter(r.FormValue("LetterTry")) {
		h.Game(strings.ToLower(r.FormValue("LetterTry")))
		details.Word = h.Word
		details.ToFind = h.ToFind
		details.Attempts = h.Attempts
		details.LetterTry = h.TriedLetters
		details.LetterKnown = h.KnownLetters
		details.Display = h.Message
		// Debug on sait jamais mdr
		// fmt.Println(details.Display)
		// fmt.Println("Lettre entrée ", r.FormValue("LetterTry"))
		// fmt.Println("Mot à trouver ", h.ToFind)
		// fmt.Println("Affichage ", h.Word)
		// fmt.Println("Lettres connues ", h.KnownLetters)
		// fmt.Println("Vous avez déjà tenté ", h.TriedLetters)
		// fmt.Println("Il vous reste ", h.Attempts, " tentatives")
		if h.Word == h.ToFind {
			details.Display = "Vous avez gagné !"
			details.Win++
		} else if h.Attempts == 0 {
			details.Display = "Vous avez perdu ! Le mot était " + h.Word
			details.Lose++
		} else {
			details.Display = "Vous devez entrer une lettre"
		}
	}
	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
