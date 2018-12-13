package hangman

import (
	"config"
	"log"
	"makeguess"
	"model"
	"newgame"
	"sync"
	"view"
	"words"
)

type Hangman struct {
	model     model.Model
	words     words.Words
	newGame   newgame.NewGame
	makeGuess makeguess.MakeGuess
	view      view.View
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

func createNewGame(config config.Config, model model.Model, words words.Words) newgame.NewGame {
	newGame, err := newgame.Create(config, model, words)
	if err != nil {
		log.Fatal(err)
	}
	return newGame
}

func createMakeGuess(config config.Config, model model.Model) makeguess.MakeGuess {
	makeGuess, err := makeguess.Create(config, model)
	if err != nil {
		log.Fatal(err)
	}
	return makeGuess
}

func createView(config config.Config, newGame newgame.NewGame, makeGuess makeguess.MakeGuess) view.View {
	view, err := view.Create(config, newGame, makeGuess)
	if err != nil {
		log.Fatal(err)
	}
	return view
}

func NewHangman(config config.Config) Hangman {
	model := createModel(config)
	words := createWords(config)
	newGame := createNewGame(config, model, words)
	makeGuess := createMakeGuess(config, model)
	view := createView(config, newGame, makeGuess)

	return Hangman{
		model:     model,
		words:     words,
		newGame:   newGame,
		makeGuess: makeGuess,
		view:      view}
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
