package words

import (
	"config"
	"errs"
)

type Words interface {
	GetWord() string
}

func NewWords(config config.Config) (Words, error) {
	switch config.WordsType() {
	case "file":
		return NewFileWord(config.WordsFile()), nil
	}
	return nil, errs.ErrWordsTypeNotSupported
}
