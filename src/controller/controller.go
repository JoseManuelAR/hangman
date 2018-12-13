package controller

import (
	"config"
	"errors"
	"model"
	"data"
	"sync"
	"word"
)

type Controller interface {
	Run(wg sync.WaitGroup) error
	GetGamesInfo() []data.GameInfo
	NewGame() (data.GameInfo, error)
	NewGuess(gameId string, guess string) (data.GameInfo, error)
}

var ErrControllerTypeNotSupported = errors.New("Controller type not supported")

func Create(config config.Config, model model.Model, words word.Words) (Controller, error) {
	switch config.ControllerType() {
	case "production":
		return NewProductionController(model, words), nil
	}
	return nil, ErrControllerTypeNotSupported
}
