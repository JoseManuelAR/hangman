package getgames

import (
	"data"
	"server/model"
)

func GetGames(model model.Model) ([]data.GameInfo, error) {
	return model.GetGamesInfo()
}
