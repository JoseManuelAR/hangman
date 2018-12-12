package controller

import (
	"log"
	"model"
	"word"
)

type productionController struct {
	model model.Model
	words word.Words
}

func NewProductionController(model model.Model, words word.Words) Controller {
	return &productionController{model: model,
		words: words}
}

func (controller productionController) GetGamesInfo() []model.GameInfo {
	return controller.model.GetGamesInfo()
}

func (controller productionController) NewGame() model.Game {
	word := controller.words.GetWord()
	log.Println(word)
	return controller.model.CreateGame(word)
}

func (controller productionController) MakeAGuess(gameId string, guess string) (*model.Game, error) {
	game, error := controller.model.GetGame(gameId)
	if error != nil {
		log.Println(error)
	}
	return game, error
	//return controller.model.MakeAGuess(gameId, guess)
}
