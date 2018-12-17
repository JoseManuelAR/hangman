package getgame

import (
	"client/remote"
	"data"
)

func GetGame(remote remote.Remote, gameId string) (data.GameInfo, error) {
	return remote.GetGame(gameId)
}
