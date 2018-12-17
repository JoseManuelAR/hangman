package remote

import (
	"config"
	"data"
	"encoding/json"
	"errs"
	"net/http"

	"github.com/go-resty/resty"
)

type remoteRest struct {
	ip   string
	port string
}

func NewRemoteRest(config config.Config) Remote {
	return &remoteRest{
		ip:   config.Ip(),
		port: config.Port()}
}

func (remote remoteRest) Start(bc chan bool) error {
	bc <- true
	return nil
}

func (remote remoteRest) NewGame() (data.GameInfo, error) {
	resp, err := resty.R().SetHeader("Content-Type", "application/json").
		Post("http://127.0.0.1:8000/hangman/v1/games")

	gameInfo := data.GameInfo{}
	if err == nil {
		json.Unmarshal([]byte(resp.String()), &gameInfo)
	}
	return gameInfo, err
}

func (remote remoteRest) NewGuess(gameId string, guess string) (data.GameInfo, error) {
	resp, err := resty.R().SetHeader("Content-Type", "application/json").
		SetBody(data.Guess{
			Guess: guess}).
		Put("http://127.0.0.1:8000/hangman/v1/games/" + gameId + "/guesses")

	gameInfo := data.GameInfo{}
	if err == nil {
		json.Unmarshal([]byte(resp.String()), &gameInfo)
	}
	return gameInfo, err
}

func (remote remoteRest) GetGame(gameId string) (data.GameInfo, error) {
	resp, err := resty.R().SetHeader("Content-Type", "application/json").
		Get("http://127.0.0.1:8000/hangman/v1/games/" + gameId)

	gameInfo := data.GameInfo{}
	if resp.StatusCode() == http.StatusOK {
		if err == nil {
			json.Unmarshal([]byte(resp.String()), &gameInfo)
		}
		return gameInfo, err
	}
	return gameInfo, errs.ErrGameNotFound
}

func (remote remoteRest) GetGames() ([]data.GameInfo, error) {
	resp, err := resty.R().SetHeader("Content-Type", "application/json").
		Get("http://127.0.0.1:8000/hangman/v1/games")

	games := make([]data.GameInfo, 0)
	if err == nil {
		json.Unmarshal([]byte(resp.String()), &games)
	}
	return games, err
}
