package word

import (
	"config"
	"errors"
)

type Words interface {
	GetWord() string
}

var ErrWordsTypeNotSupported = errors.New("Words type not supported")

func Create(config config.Config) (Words, error) {
	switch config.WordsType() {
	case "file":
		return NewFileWord(config.WordsFile()), nil
	}
	return nil, ErrWordsTypeNotSupported
}
