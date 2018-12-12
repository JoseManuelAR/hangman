package model

import (
	"github.com/google/uuid"
	"strings"
)

type Game struct {
	Id      string
	Letters []string
	Used    map[string]bool
}

type GameInfo struct {
	Id string
	RevealedWord string
}

func newGame(word string) Game {
	letters := strings.Split(word, "")
	return Game{
		Id:      uuid.New().String(),
		Letters: letters,
		Used: make(map[string]bool)}
}

func newGameInfo(game Game) GameInfo {
	return GameInfo{
		Id: game.Id,
		RevealedWord: revealWord(game.Letters, game.Used)}
}

func revealWord(letters []string, used map[string]bool) string {
	revealedWord := ""

	for _, wordLetter := range letters {
		if used[wordLetter] {
			revealedWord += wordLetter
		} else {
			revealedWord += "_"
		}
	}
	return revealedWord
}
