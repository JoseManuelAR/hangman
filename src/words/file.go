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
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	words := strings.Split(str, "\n")
	return words
}
