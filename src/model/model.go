package model

import (
	"config"
	"data"
	"errs"
	"sync"
)

type Model interface {
	Run(wg sync.WaitGroup) error
	CreateGame(word string) data.Game
	UpdateGame(id string, game data.Game) error
	GetGame(id string) (data.Game, error)
	GetGamesInfo() []data.GameInfo
}

func Create(config config.Config) (Model, error) {
	switch config.ModelType() {
	case "memory":
		return NewMemoryModel(), nil
	}
	return nil, errs.ErrModelTypeNotSupported
}
