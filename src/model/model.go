package model

import (
	"config"
	"errors"
	"data"
)

type Model interface {
	Run() error
	CreateGame(word string) data.Game
	UpdateGame(id string, game data.Game) error
	GetGame(id string) (data.Game, error)
	GetGamesInfo() []data.GameInfo
}

var ErrModelTypeNotSupported = errors.New("Model type not supported")

func Create(config config.Config) (Model, error) {
	switch config.ModelType() {
	case "memory":
		return NewMemoryModel(), nil
	}
	return nil, ErrModelTypeNotSupported
}
