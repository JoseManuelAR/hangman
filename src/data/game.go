package data

import (
	"strconv"
	"strings"
)

type Game struct {
	Id        string
	Letters   []string
	Used      map[string]bool
	Status    GameStatus
	TurnsLeft int
}

type GameInfo struct {
	Id           string
	RevealedWord string
	Status       string
	TurnsLeft    string
}

func NewGame(id string, word string) Game {
	letters := strings.Split(word, "")
	return Game{
		Id:        id,
		Letters:   letters,
		Used:      make(map[string]bool),
		Status:    Initial,
		TurnsLeft: 5}
}

func NewGameInfo(game Game) GameInfo {
	return GameInfo{
		Id:           game.Id,
		RevealedWord: revealWord(game.Letters, game.Used),
		Status:       game.Status.String(),
		TurnsLeft:    strconv.Itoa(game.TurnsLeft)}
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
