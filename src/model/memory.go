package model

import (
	"errors"
	"log"
	"sync"
)

type memoryModel struct {
	games map[string]Game
}

func NewMemoryModel() Model {
	return memoryModel{games: make(map[string]Game)}
}

func (model memoryModel) Start(wg sync.WaitGroup) error {
	log.Println("Starting memory model...")
	//words := readWordsFromFile(model.words_file)
	return nil
}

func (model memoryModel) CreateGame(word string) Game {
	game := newGame(word)
	model.games[game.Id] = game
	return model.games[game.Id]
}

func (model memoryModel) GetGame(id string) (*Game, error) {
	game, ok := model.games[id]
	if !ok {
		return nil, errors.New("game " + id + " not found")
	}
	return &game, nil
}

func (model memoryModel) GetGamesInfo() []GameInfo {
	gamesInfo := make([]GameInfo, 0, len(model.games))
	for _, game := range model.games {
		gamesInfo = append(gamesInfo, newGameInfo(game))
	}
	return gamesInfo
}
