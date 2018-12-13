package model

import (
	"data"
	"err"
	uuid "github.com/google/uuid"
	"log"
	"sync"
)

type memoryModel struct {
	games map[string]data.Game
}

func NewMemoryModel() Model {
	return memoryModel{games: make(map[string]data.Game)}
}

func (model memoryModel) Run(wg sync.WaitGroup) error {
	log.Println("Starting memory model...")
	//words := readWordsFromFile(model.words_file)
	return nil
}

func (model memoryModel) CreateGame(word string) data.Game {
	id := uuid.New().String()
	model.games[id] = data.NewGame(id, word)
	return model.games[id]
}

func (model memoryModel) UpdateGame(id string, game data.Game) error {
	model.games[id] = game
	return nil
}

func (model memoryModel) GetGame(id string) (data.Game, error) {
	game, ok := model.games[id]
	if !ok {
		return data.Game{}, err.ErrGameNotFound
	}
	return game, nil
}

func (model memoryModel) GetGamesInfo() []data.GameInfo {
	gamesInfo := make([]data.GameInfo, 0, len(model.games))
	for _, game := range model.games {
		gamesInfo = append(gamesInfo, data.NewGameInfo(game))
	}
	return gamesInfo
}
