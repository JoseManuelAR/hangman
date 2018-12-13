package hangman

import (
	"config"
	"controller"
	"log"
	"model"
	"sync"
	"view"
	"word"
)

type Hangman struct {
	view       view.View
	controller controller.Controller
}

func createModel(config config.Config) model.Model {
	model, err := model.Create(config)
	if err != nil {
		log.Fatal(err)
	}
	return model
}

func createWord(config config.Config) word.Words {
	word, err := word.Create(config)
	if err != nil {
		log.Fatal(err)
	}
	return word
}

func NewHangman(config config.Config) Hangman {
	controller, err := controller.Create(config, createModel(config), createWord(config))
	if err != nil {
		log.Fatal(err)
	}
	view, err := view.Create(config, controller)
	if err != nil {
		log.Fatal(err)
	}
	return Hangman{view: view,
		controller: controller}
}

func (hangman Hangman) Run() error {
	log.Println("Starting hangman...")
	var wg sync.WaitGroup
	wg.Add(2)
	go hangman.controller.Run(wg)
	go hangman.view.Run(wg)
	wg.Wait()
	return nil
}
