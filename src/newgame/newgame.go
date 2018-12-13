package newgame

import (
	"config"
	"data"
	"err"
	"model"
	"words"
)

type NewGame interface {
	Execute() (data.GameInfo, error)
}

type productionNewGame struct {
	model model.Model
	words words.Words
}

func (newGame productionNewGame) Execute() (data.GameInfo, error) {
	word := newGame.words.GetWord()
	return data.NewGameInfo(newGame.model.CreateGame(word)), nil
}

func Create(config config.Config, model model.Model, words words.Words) (NewGame, error) {
	switch config.ControllerType() {
	case "production":
		return &productionNewGame{model: model,
			words: words}, nil
	}
	return nil, err.ErrControllerTypeNotSupported
}
