package controller

import (
	"model"
)

type Controller interface {
	GetGamesInfo() []model.GameInfo
	NewGame() model.Game
	MakeAGuess(gameId string, guess string) (*model.Game, error)
}
