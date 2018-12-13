package controller

import (
	"config"
	"data"
	"err"
)

type Controller interface {
	GetGamesInfo() []data.GameInfo
	//NewGuess(gameId string, guess string) (data.GameInfo, error)
}

func Create(config config.Config) (Controller, error) {
	// switch config.ControllerType() {
	// case "production":
	// 	return NewProductionController(model, words), nil
	// }
	return nil, err.ErrControllerTypeNotSupported
}
