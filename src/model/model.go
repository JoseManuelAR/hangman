package model

import (
	"sync"
)

type Model interface {
	Start(wg sync.WaitGroup) error
	CreateGame(word string) Game
	GetGame(id string) (*Game, error)
	GetGamesInfo() []GameInfo
}
