package newgame

import (
	"client/remote"
	"data"
)

func NewGame(remote remote.Remote) (data.GameInfo, error) {
	return remote.NewGame()
}
