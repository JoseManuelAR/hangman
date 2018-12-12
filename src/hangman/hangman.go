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
	model      model.Model
	controller controller.Controller
}

func createModel(name string) model.Model {
	switch name {
	case "memory":
		return model.NewMemoryModel()
	}
	return nil
}

func createWords(name string, file string) word.Words {
	switch name {
	case "file":
		return word.NewFileWord(file)
	}
	return nil
}

func createController(name string, model model.Model, words word.Words) controller.Controller {
	switch name {
	case "production":
		return controller.NewProductionController(model, words)
	}
	return nil
}

func createView(name string, ip string, port string, controller controller.Controller) view.View {
	switch name {
	case "rest":
		return view.NewRestServer(controller, view.RestConfig{ip, port})
	}
	return nil
}

func NewHangman(config config.Config) Hangman {
	model := createModel(config.ModelType)
	if model == nil {
		log.Fatal("Error creating model. Exiting program")
	}
	words := createWords(config.WordsType, config.WordsFile)
	if words == nil {
		log.Fatal("Error creating words generator. Exiting program")
	}
	controller := createController(config.ControllerType, model, words)
	if controller == nil {
		log.Fatal("Error creating contoller. Exiting program")
	}
	view := createView(config.ViewType, config.Ip, config.Port, controller)
	if view == nil {
		log.Fatal("Error creating view. Exiting program")
	}
	return Hangman{view: view,
		model: model}
}

func (hangman Hangman) Start() error {
	log.Println("Starting hangman...")

	var wg sync.WaitGroup
	wg.Add(2)
	go hangman.model.Start(wg)
	go hangman.view.Start(wg)
	wg.Wait()
	return nil
}
