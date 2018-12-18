package data

import (
	"strings"
)

type Game struct {
	Id        string
	Letters   []string
	Used      map[string]bool
	Status    GameStatus
	TurnsLeft int
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
