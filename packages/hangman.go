package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var h *HangManData

type HangManData struct {
	Word         string
	ToFind       string
	Attempts     int
	KnownLetters []string
	TriedLetters []string
	Message      string
}

func main() {
	path := os.Args[1]
	h.Attempts = 10
	h.ToFind = RandomWord(path)
	n := (len(h.ToFind) / 2) - 1
	h.KnownLetters = append(h.KnownLetters, string(h.ToFind[n]))
	h.Updateword()
}

func (h *HangManData) Init(path string) { //func to initialize the game
	h.Attempts = 10
	h.KnownLetters = nil
	h.TriedLetters = nil
	h.Message = ""
	h.ToFind = RandomWord(path)
	n := (len(h.ToFind) / 2) - 1
	h.KnownLetters = append(h.KnownLetters, string(h.ToFind[n]))
	h.Updateword()
}

func (h *HangManData) Game(entry string) { //func to play the game
	if h.Word == h.ToFind {
		h.Message = "Vous avez gagné !"
	} else if h.Attempts == 0 {
		h.Message = "Vous avez perdu !"
	} else {
		if len(entry) == 1 {
			if Isintheword(h.ToFind, entry) {
				if AlreadyKnown(h, entry) {
					h.Message = "Vous avez déjà trouvé cette lettre"
				} else {
					h.KnownLetters = append(h.KnownLetters, string(entry))
					h.Updateword()
					h.Message = "Vous avez trouvé une lettre, Bien joué !"
				}
			} else if !Alreadytried(h, entry) {
				h.Attempts--
				h.TriedLetters = append(h.TriedLetters, entry)
				h.Message = "Cette lettre n'est pas dans le mot, essayez encore !, vous perdez un point"
			} else {
				h.Message = "Vous avez déjà essayé cette lettre"
			}
		} else if len(entry) >= 2 {
			if entry == h.ToFind {
				h.Word = h.ToFind
				h.Message = "Vous avez trouvé le mot, Bravo !"
			} else {
				h.Message = "Ce n'est pas le mot, essayez encore !, Vous perdez deux points"
				h.Attempts--
				h.Attempts--
			}
		}
	}
}
func Isintheword(word string, letter string) bool { //func to check if the letter is in the word
	for _, v := range word {
		if string(v) == letter {
			return true
		}
	}
	return false
}
func Alreadytried(h *HangManData, letter string) bool { //func to check if the letter is already tried
	for _, v := range h.TriedLetters {
		if v == letter {
			return true
		}
	}
	return false
}
func AlreadyKnown(h *HangManData, letter string) bool { //func to check if the letter is already known
	for _, v := range h.KnownLetters {
		if v == letter {
			return true
		}
	}
	return false
}
func (h *HangManData) Updateword() { //func to update the display
	h.Word = ""
	for _, v := range h.ToFind {
		if AlreadyKnown(h, string(v)) {
			h.Word += string(v)
		} else {
			h.Word += "_"
		}
	}
}
func RandomWord(arg string) string { //func to get a random word from the file.txt
	//read the file
	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//make a slice of words
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	//return a random word
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}
