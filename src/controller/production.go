package controller

import (
	"data"
	"model"
	"words"
)

type productionController struct {
	model model.Model
	words words.Words
}

func NewProductionController(model model.Model, words words.Words) Controller {
	return &productionController{model: model,
		words: words}
}

func (controller productionController) GetGamesInfo() []data.GameInfo {
	return controller.model.GetGamesInfo()
}

// func (controller productionController) NewGame() (data.GameInfo, error) {
// 	word := controller.words.GetWord()
// 	return data.NewGameInfo(controller.model.CreateGame(word)), nil
// }

// func (controller productionController) NewGuess(gameId string, guess string) (data.GameInfo, error) {
// 	game, err := controller.model.GetGame(gameId)
// 	if err != nil {
// 		return data.GameInfo{}, err
// 	}
// 	if game.Status == data.Lost || game.Status == data.Won {
// 		return data.NewGameInfo(game), nil
// 	}
// 	if game.Used[guess] {
// 		game.Status = data.AlreadyGuessed
// 	} else if letterInWord(guess, game.Letters) {
// 		game.Used[guess] = true
// 		game.Status = data.GoodGuess
// 		if hasWon(game.Letters, game.Used) {
// 			game.Status = data.Won
// 		}
// 	} else {
// 		game.TurnsLeft--
// 		game.Status = data.BadGuess
// 		game.Used[guess] = true
// 		if game.TurnsLeft == 0 {
// 			game.Status = data.Lost
// 		}
// 	}
// 	err = controller.model.UpdateGame(gameId, game)
// 	return data.NewGameInfo(game), err
// }
