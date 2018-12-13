package view

import (
	"config"
	"controller"
	"errors"
	"sync"
)

var ErrViewTypeNotSupported = errors.New("View type not supported")

type View interface {
	Run(wg sync.WaitGroup) error
}

func Create(config config.Config, controller controller.Controller) (View, error) {
	switch config.ViewType() {
	case "rest":
		return NewRestServer(config, controller), nil
	}
	return nil, ErrViewTypeNotSupported
}
