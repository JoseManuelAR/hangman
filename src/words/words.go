package words

import (
	config "config/server"
	"errs"
)

type Words interface {
	GetWord() string
}

func Create(config config.Config) (Words, error) {
	switch config.WordsType() {
	case "file":
		return NewFileWord(config.WordsFile()), nil
	}
	return nil, errs.ErrWordsTypeNotSupported
}
