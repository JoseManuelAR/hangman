package model

import (
	"config"
	"data"
	"errors"
	"sync"
)

type Model interface {
	Run(wg sync.WaitGroup) error
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
