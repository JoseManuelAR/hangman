package getgame

import (
	"data"
	"server/model"
)

func GetGame(model model.Model, gameId string) (data.GameInfo, error) {
	return model.GetGameInfo(gameId)
}
