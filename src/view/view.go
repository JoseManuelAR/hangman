package view

import (
	config "config/server"
	"controller"
	"errs"
	"sync"
)

type View interface {
	Run(wg sync.WaitGroup) error
}

func Create(config config.Config, controller controller.Controller) (View, error) {
	switch config.ViewType() {
	case "rest":
		return NewRestServer(config, controller), nil
	}
	return nil, errs.ErrViewTypeNotSupported
}
