package getgames

import (
	"data"
	"model"
)

func GetGames(model model.Model) ([]data.GameInfo, error) {
	return model.GetGamesInfo(), nil
}
