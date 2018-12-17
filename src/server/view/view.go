package view

import (
	"config"
	"errs"
	"server/controller"
)

type View interface {
	Start(bc chan bool) error
}

func NewView(config config.Config, controller controller.Controller) (View, error) {
	switch config.RemoteType() {
	case "rest":
		return NewRestServer(config, controller), nil
	}
	return nil, errs.ErrViewTypeNotSupported
}
