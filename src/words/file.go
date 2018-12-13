package words

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type fileWords struct {
	words []string
}

func NewFileWord(filePath string) Words {
	words := readWordsFromFile(filePath)
	log.Println("Number of words:", len(words))
	return fileWords{words: words}
}

func (words fileWords) GetWord() string {
	wordIndex := getSystemRandomInt(len(words.words))
	return words.words[wordIndex]
}

func getSystemRandomInt(i int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}

func readWordsFromFile(filePath string) []string {
	b, err := ioutil.ReadFile(filePath) // read words from file
	if err != nil {
		log.Fatal(err)
	}
	str := string(b) // convert content to a 'string'
	words := strings.Split(str, "\n")
	return words
}
