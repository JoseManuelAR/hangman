package view

import (
	"client/controller"
	"config"
	"errs"
)

type View interface {
	Start(bc chan bool) error
}

func NewView(config config.Config, controller controller.Controller) (View, error) {
	switch config.ViewType() {
	case "cli":
		return NewCliView(controller), nil
	}
	return nil, errs.ErrViewTypeNotSupported
}
