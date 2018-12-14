package errs

import (
	"errors"
)

var ErrModelTypeNotSupported = errors.New("Model type not supported")
var ErrViewTypeNotSupported = errors.New("View type not supported")
var ErrWordsTypeNotSupported = errors.New("Words type not supported")
var ErrEmptyGuess = errors.New("Empty guess")
var ErrGameNotFound = errors.New("Game not found")
