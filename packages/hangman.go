package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word         string
	ToFind       string
	Attempts     int
	KnownLetters []string
	TriedLetters []string
}

func main() {
	h := HangManData{}
	h.Init()
}

func (h *HangManData) Init() { //func to initialize the game
	path := os.Args[1]
	h.Attempts = 10
	h.ToFind = RandomWord(path)
	// for _, v := range h.ToFind {
	// 	h.Word += "_"
	// 	_ = v //to avoid the error
	// }
	n := (len(h.ToFind) / 2) - 1
	h.KnownLetters = append(h.KnownLetters, string(h.ToFind[n]))
	h.Updateword()
}

func (h *HangManData) Game(entry string) int { //func to play the game
	if len(entry) == 1 {
		if Isintheword(h.ToFind, entry) {
			if AlreadyKnown(h, entry) {
				return 1
			} else {
				h.KnownLetters = append(h.KnownLetters, entry)
				h.Updateword()
				// fmt.Println("Bien joué !")
				return 0
			}
		} else if !Alreadytried(h, entry) {
			h.Attempts--
			// fmt.Println("La lettre ", entry, " n'est pas dans le mot")
			h.TriedLetters = append(h.TriedLetters, entry)
			return 2
		} else {
			// fmt.Println("Tu as déjà essayé cette lettre")
			return 3
		}
	} else if len(entry) >= 2 {
		if entry == h.ToFind {
			return 4
		} else {
			h.Attempts--
			h.Attempts--
			return 5
		}
	}
	return 99
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
