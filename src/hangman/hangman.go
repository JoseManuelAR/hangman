package hangman

import (
	"config"
	"controller"
	"log"
	"model"
	"sync"
	"view"
	"words"
)

type Hangman struct {
	model model.Model
	view  view.View
}

func NewHangman(config config.Config) Hangman {
	model := createModel(config)
	words := createWords(config)
	controller := createController(model, words)
	view := createView(config, controller)

	return Hangman{
		model: model,
		view:  view}
}

func (hangman Hangman) Run() error {
	log.Println("Starting hangman...")
	var wg sync.WaitGroup
	wg.Add(2)
	go hangman.model.Run(wg)
	go hangman.view.Run(wg)
	wg.Wait()
	return nil
}

func createModel(config config.Config) model.Model {
	model, err := model.Create(config)
	if err != nil {
		log.Fatal(err)
	}
	return model
}

func createWords(config config.Config) words.Words {
	word, err := words.Create(config)
	if err != nil {
		log.Fatal(err)
	}
	return word
}

func createController(model model.Model, words words.Words) controller.Controller {
	controller, err := controller.Create(model, words)
	if err != nil {
		log.Fatal(err)
	}
	return controller
}

func createView(config config.Config, controller controller.Controller) view.View {
	view, err := view.Create(config, controller)
	if err != nil {
		log.Fatal(err)
	}
	return view
}
