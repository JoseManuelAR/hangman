package err

import (
	"errors"
)

var ErrControllerTypeNotSupported = errors.New("Controller type not supported")
var ErrGameNotFound = errors.New("Game not found")
