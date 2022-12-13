package main

import (
	"fmt"
	"html/template"
	"net/http"
	hangman "serv-hangman/packages"
	"strconv"
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
	Img         string
	Winrate     float64
	GameEnd     string
}

var h hangman.HangManData

var details = User{
	Username:   "",
	Difficulty: "",
	Attempts:   11,
	GameEnd:    "none",
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion img
	img := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", img))
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
	if details.Attempts == 11 || details.Attempts <= 0 || details.Word == details.ToFind {
		h = hangman.HangManData{}
		details.GameEnd = "none"
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
		if h.Word == h.ToFind {
			details.Win++
			h.ToFind = h.Word
			details.GameEnd = "visible"
		} else if h.Attempts <= 0 {
			h.Attempts = 0
			details.Lose++
			details.GameEnd = "visible"
		} else {
			details.Display = "Vous devez entrer une lettre"
		}
	}
	details.Img = "/img/" + strconv.Itoa(h.Attempts) + ".png"
	if details.Win != 0 || details.Lose != 0 {
		details.Winrate = float64(details.Win) / float64(details.Win+details.Lose) * 100
	}
	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
