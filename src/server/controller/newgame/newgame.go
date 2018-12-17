package newgame

import (
	"data"
	"server/model"
	"server/words"
)

func NewGame(model model.Model, words words.Words) (data.GameInfo, error) {
	word := words.GetWord()
	return data.NewGameInfo(model.CreateGame(word)), nil
}
