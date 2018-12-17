package getgames

import (
	"client/remote"
	"data"
)

func GetGames(remote remote.Remote) ([]data.GameInfo, error) {
	return remote.GetGames()
}
