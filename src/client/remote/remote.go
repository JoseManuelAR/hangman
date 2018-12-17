package remote

import (
	"config"
	"data"
	"errs"
)

type Remote interface {
	Start(bc chan bool) error
	NewGame() (data.GameInfo, error)
	NewGuess(gameId string, guess string) (data.GameInfo, error)
	GetGame(gameId string) (data.GameInfo, error)
	GetGames() ([]data.GameInfo, error)
}

func NewRemote(config config.Config) (Remote, error) {
	switch config.RemoteType() {
	case "rest":
		return NewRemoteRest(config), nil
	}
	return nil, errs.ErrModelTypeNotSupported
}
