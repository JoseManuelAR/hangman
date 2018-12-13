package view

import (
	"config"
	"errors"
	"makeguess"
	"newgame"
	"sync"
)

var ErrViewTypeNotSupported = errors.New("View type not supported")

type View interface {
	Run(wg sync.WaitGroup) error
}

func Create(config config.Config, newGame newgame.NewGame, makeGuess makeguess.MakeGuess) (View, error) {
	switch config.ViewType() {
	case "rest":
		return NewRestServer(config, newGame, makeGuess), nil
	}
	return nil, ErrViewTypeNotSupported
}
