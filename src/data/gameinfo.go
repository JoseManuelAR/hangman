package data

import (
	"strconv"
)

type GameInfo struct {
	Id           string
	RevealedWord string
	Status       string
	TurnsLeft    string
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
