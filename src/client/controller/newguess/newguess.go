package newguess

import (
	"client/remote"
	"data"
)

func NewGuess(remote remote.Remote, gameId string, guess string) (data.GameInfo, error) {
	return remote.NewGuess(gameId, guess)
}
