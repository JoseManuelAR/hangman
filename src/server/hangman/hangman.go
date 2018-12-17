package hangman

import (
	"config"
	"log"
	"server/controller"
	"server/view"
)

type Hangman struct {
	controller controller.Controller
	view       view.View
}

func NewHangman(config config.Config) Hangman {
	controller, err := controller.NewController(config)
	if err != nil {
		log.Fatal(err)
	}
	view, err := view.NewView(config, controller)
	if err != nil {
		log.Fatal(err)
	}
	return Hangman{
		controller: controller,
		view:       view}
}

func (hangman Hangman) Start() error {
	log.Println("Starting hangman...")
	done := make(chan bool)
	go hangman.controller.Start(done)
	<-done
	go hangman.view.Start(done)
	<-done
	return nil
}
