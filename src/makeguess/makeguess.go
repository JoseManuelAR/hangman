package makeguess

import (
	"config"
	"data"
	"err"
	"model"
)

type MakeGuess interface {
	Execute(gameId string, guess string) (data.GameInfo, error)
}

type productionMakeGuess struct {
	model model.Model
}

func (makeGuess productionMakeGuess) Execute(gameId string, guess string) (data.GameInfo, error) {
	game, err := makeGuess.model.GetGame(gameId)
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
	err = makeGuess.model.UpdateGame(gameId, game)
	return data.NewGameInfo(game), err
}

func Create(config config.Config, model model.Model) (MakeGuess, error) {
	switch config.ControllerType() {
	case "production":
		return &productionMakeGuess{model: model}, nil
	}
	return nil, err.ErrControllerTypeNotSupported
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
