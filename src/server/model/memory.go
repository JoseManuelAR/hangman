package model

import (
	"data"
	"errs"
	"log"
	"sync"

	"github.com/google/uuid"
)

type memoryModel struct {
	mutex sync.RWMutex
	games map[string]data.Game
}

func NewMemoryModel() Model {
	return memoryModel{games: make(map[string]data.Game)}
}

func (model memoryModel) Start(bc chan bool) error {
	log.Println("Starting memory model...")
	bc <- true
	return nil
}

func (model memoryModel) CreateGame(word string) data.Game {
	model.mutex.Lock()
	defer model.mutex.Unlock()
	id := uuid.New().String()
	model.games[id] = data.NewGame(id, word)
	return model.games[id]
}

func (model memoryModel) UpdateGame(id string, game data.Game) error {
	model.mutex.Lock()
	defer model.mutex.Unlock()
	model.games[id] = game
	return nil
}

func (model memoryModel) GetGame(id string) (data.Game, error) {
	model.mutex.RLock()
	defer model.mutex.RUnlock()
	game, ok := model.games[id]
	if !ok {
		return data.Game{}, errs.ErrGameNotFound
	}
	return game, nil
}

func (model memoryModel) GetGameInfo(id string) (data.GameInfo, error) {
	model.mutex.RLock()
	defer model.mutex.RUnlock()
	game, ok := model.games[id]
	if !ok {
		return data.GameInfo{}, errs.ErrGameNotFound
	}
	return data.NewGameInfo(game), nil
}

func (model memoryModel) GetGamesInfo() ([]data.GameInfo, error) {
	model.mutex.RLock()
	defer model.mutex.RUnlock()
	gamesInfo := make([]data.GameInfo, 0, len(model.games))
	for _, game := range model.games {
		gamesInfo = append(gamesInfo, data.NewGameInfo(game))
	}
	return gamesInfo, nil
}
