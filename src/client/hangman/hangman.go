package hangman

import (
	"client/controller"
	"client/view"
	"config"
	"fmt"
)

type Hangman struct {
	controller controller.Controller
	view       view.View
}

func NewHangman(config config.Config) Hangman {
	controller, err := controller.NewController(config)
	if err != nil {
		fmt.Println(err)
	}
	view, err := view.NewView(config, controller)
	if err != nil {
		fmt.Println(err)
	}
	return Hangman{
		controller: controller,
		view:       view}
}

func (hangman Hangman) Start() error {
	done := make(chan bool)
	go hangman.controller.Start(done)
	<-done
	go hangman.view.Start(done)
	<-done
	return nil
}
