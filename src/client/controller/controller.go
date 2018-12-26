package controller

import (
	"client/controller/getgame"
	"client/controller/getgames"
	"client/controller/newgame"
	"client/controller/newguess"
	"client/remote"
	"config"
	"data"
	"log"
)

type Controller struct {
	remote remote.Remote
}

func NewController(config config.Config) (Controller, error) {
	remote, err := remote.NewRemote(config)
	if err != nil {
		return Controller{}, err
	}
	return Controller{remote: remote}, nil
}

func (controller Controller) Start(bc chan bool) error {
	log.Println("Starting controller...")
	controller.remote.Start(bc)
	return nil
}

func (controller Controller) NewGame() (data.GameInfo, error) {
	return newgame.NewGame(controller.remote)
}

func (controller Controller) GetGame(gameId string) (data.GameInfo, error) {
	return getgame.GetGame(controller.remote, gameId)
}

func (controller Controller) GetGames() ([]data.GameInfo, error) {
	return getgames.GetGames(controller.remote)
}

func (controller Controller) NewGuess(gameId string, guess string) (data.GameInfo, error) {
	return newguess.NewGuess(controller.remote, gameId, guess)
}
