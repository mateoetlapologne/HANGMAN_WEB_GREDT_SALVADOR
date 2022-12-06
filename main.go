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
	h := hangman.HangManData{}
	h.Init()
	details.Word = h.ToFind
	details.WordDisplay = h.Word

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
	h.Game(r.FormValue("LetterTry"))
	details.TryNumber--
	//details.LetterTry = append(details.LetterTry, r.FormValue("LetterTry"))
	//debug pour voir les valeurs
	fmt.Println("Lettre deja tentée", details.LetterTry)
	fmt.Println("Lettre entrée ", r.FormValue("LetterTry"))
	fmt.Println("Mot à trouver ", details.Word)
	fmt.Println("Affichage ", details.WordDisplay)
	fmt.Println("Il vous reste ", h.Attempts, " tentatives")
	if h.Word == h.ToFind || h.Attempts == 0 {
		fmt.Println("Vous avez gagné !")
	}
	if h.Game(r.FormValue("LetterTry")) == 1 {
		fmt.Println("Vous avez déjà essayé cette lettre")
	}
	if h.Game(r.FormValue("LetterTry")) == 0 {
		fmt.Println("Bien joué !")
	}
	if h.Game(r.FormValue("LetterTry")) == 2 {
		fmt.Println("La lettre ", r.FormValue("LetterTry"), " n'est pas dans le mot")
	}
	if h.Game(r.FormValue("LetterTry")) == 3 {
		fmt.Println("Vous avez perdu !")
	}
	if h.Game(r.FormValue("LetterTry")) == 4 {
		fmt.Println("Vous avez gagné !")
	}
	if h.Word == h.ToFind {
		fmt.Println("Vous avez gagné !")
	}
	//gestion html
	tmpl1 := template.Must(template.ParseFiles("templates/game.html"))
	tmpl1.Execute(w, details)
}
