package controller

import (
	"data"
	"getgames"
	"model"
	"newgame"
	"newguess"
	"words"
)

type Controller struct {
	model model.Model
	words words.Words
}

func (controller Controller) NewGame() (data.GameInfo, error) {
	return newgame.NewGame(controller.model, controller.words)
}

func (controller Controller) GetGames() ([]data.GameInfo, error) {
	return getgames.GetGames(controller.model)
}

func (controller Controller) NewGuess(gameId string, guess string) (data.GameInfo, error) {
	return newguess.NewGuess(controller.model, gameId, guess)
}

func Create(model model.Model, words words.Words) (Controller, error) {
	return Controller{
		model: model,
		words: words}, nil
}
