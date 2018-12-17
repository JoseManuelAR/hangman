package model

import (
	"config"
	"data"
	"errs"
)

type Model interface {
	Start(chan bool) error
	CreateGame(word string) data.Game
	UpdateGame(id string, game data.Game) error
	GetGame(id string) (data.Game, error)
	GetGameInfo(id string) (data.GameInfo, error)
	GetGamesInfo() ([]data.GameInfo, error)
}

func NewModel(config config.Config) (Model, error) {
	switch config.ModelType() {
	case "memory":
		return NewMemoryModel(), nil
	}
	return nil, errs.ErrModelTypeNotSupported
}
