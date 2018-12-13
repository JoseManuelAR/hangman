package controller

import (
	"model"
	"sync"
	"word"
	"data"
)

type productionController struct {
	model model.Model
	words word.Words
}

func letterInWord(guess string, letters []string) bool {
	for _, letter := range letters {
		if guess == letter {
			return true
		}
	}
	return false
}

func hasWon(letters []string, used map[string]bool) bool {
	for _, letter := range letters {
		if !used[letter] {
			return false
		}
	}
	return true
}

func NewProductionController(model model.Model, words word.Words) Controller {
	return &productionController{model: model,
		words: words}
}

func (controller productionController) Run(wg sync.WaitGroup) error {
	controller.model.Run()
	return nil
}

func (controller productionController) GetGamesInfo() []data.GameInfo {
	return controller.model.GetGamesInfo()
}

func (controller productionController) NewGame() (data.GameInfo, error) {
	word := controller.words.GetWord()
	return data.NewGameInfo(controller.model.CreateGame(word)), nil
}

func (controller productionController) NewGuess(gameId string, guess string) (data.GameInfo, error) {
	game, err := controller.model.GetGame(gameId)
	if err != nil {
		return data.GameInfo{}, err
	}
	if game.Status == data.Lost || game.Status == data.Won {
		return data.NewGameInfo(game), nil
	}
	if game.Used[guess] {
		game.Status = data.AlreadyGuessed
	} else if letterInWord(guess, game.Letters) {
		game.Used[guess] = true
		game.Status = data.GoodGuess
		if hasWon(game.Letters, game.Used) {
			game.Status = data.Won
		}
	} else {
		game.TurnsLeft--
		game.Status = data.BadGuess
		game.Used[guess] = true
		if game.TurnsLeft == 0 {
			game.Status = data.Lost
		}
	}
	err = controller.model.UpdateGame(gameId, game)
	return data.NewGameInfo(game), err
}
