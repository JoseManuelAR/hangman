package newguess

import (
	"data"
	"errs"
	"server/model"
)

func NewGuess(model model.Model, gameId string, guess string) (data.GameInfo, error) {
	if len(guess) == 0 {
		return data.GameInfo{}, errs.ErrEmptyGuess
	}
	game, err := model.GetGame(gameId)
	if err != nil {
		return data.GameInfo{}, err
	}
	if game.Status == data.LostCode || game.Status == data.WonCode {
		return data.NewGameInfo(game), nil
	}
	if game.Used[guess] {
		game.Status = data.AlreadyGuessedCode
	} else if letterInWord(guess, game.Letters) {
		game.Used[guess] = true
		game.Status = data.GoodGuessCode
		if hasWon(game.Letters, game.Used) {
			game.Status = data.WonCode
		}
	} else {
		game.TurnsLeft--
		game.Status = data.BadGuessCode
		game.Used[guess] = true
		if game.TurnsLeft == 0 {
			game.Status = data.LostCode
		}
	}
	err = model.UpdateGame(gameId, game)
	return data.NewGameInfo(game), err
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
