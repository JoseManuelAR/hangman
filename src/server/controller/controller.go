package controller

import (
	"config"
	"data"
	"log"
	"server/controller/getgame"
	"server/controller/getgames"
	"server/controller/newgame"
	"server/controller/newguess"
	"server/model"
	"server/words"
)

type Controller struct {
	model model.Model
	words words.Words
}

func NewController(config config.Config) (Controller, error) {
	model, err := model.NewModel(config)
	if err != nil {
		return Controller{}, err
	}
	words, err := words.NewWords(config)
	if err != nil {
		return Controller{}, err
	}
	return Controller{model: model,
		words: words}, nil
}

func (controller Controller) Start(bc chan bool) error {
	log.Println("Starting controller...")
	controller.model.Start(bc)
	return nil
}

func (controller Controller) NewGame() (data.GameInfo, error) {
	return newgame.NewGame(controller.model, controller.words)
}

func (controller Controller) GetGame(gameId string) (data.GameInfo, error) {
	return getgame.GetGame(controller.model, gameId)
}

func (controller Controller) GetGames() ([]data.GameInfo, error) {
	return getgames.GetGames(controller.model)
}

func (controller Controller) NewGuess(gameId string, guess string) (data.GameInfo, error) {
	return newguess.NewGuess(controller.model, gameId, guess)
}
