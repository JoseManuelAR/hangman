package controller

import (
	"mocks"
	"testing"
	"data"
	"error"
	gomock "github.com/golang/mock/gomock"
)

func TestLetterInWord(t *testing.T) {
	word := []string{"t", "e", "s", "t"}
	guess := "e"
	hasLetter := letterInWord(guess, word)
	if hasLetter != true {
		t.Errorf("Word %s does not contain letter %s", word, guess)
	}
}

func TestNotLetterInWord(t *testing.T) {
	word := []string{"t", "e", "s", "t"}
	guess := "c"
	hasLetter := letterInWord(guess, word)
	if hasLetter == true {
		t.Errorf("Word %s should not contain letter %s", word, guess)
	}
}

func TestGuessToNonExistentGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	words := mocks.NewMockWords(mockCtrl)

	model.EXPECT().GetGame("XXXX").Return(data.Game{}, errors.ErrGameNotFound)
	c := NewProductionController(model, words)
	_, err := c.NewGuess("XXXX", "a")

	if err != errors.ErrGameNotFound {
		t.Fail()
	}
}

func TestGuessToLostGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	words := mocks.NewMockWords(mockCtrl)

	model.EXPECT().GetGame("XXXX").Return(data.Game{Status: data.Lost}, nil)
	c := NewProductionController(model, words)
	game, err := c.NewGuess("XXXX", "a")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.Lost.String() {
		t.Fail()
	}
}

func TestGuessToWonGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	words := mocks.NewMockWords(mockCtrl)

	model.EXPECT().GetGame("XXXX").Return(data.Game{Status: data.Won}, nil)
	c := NewProductionController(model, words)
	game, err := c.NewGuess("XXXX", "a")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.Won.String() {
		t.Fail()
	}
}

func TestGuessAlreadyGuessed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	words := mocks.NewMockWords(mockCtrl)

	used := make(map[string]bool)
	used["a"] = true
	model.EXPECT().GetGame("XXXX").Return(data.Game{Status: data.BadGuess, Used: used}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	c := NewProductionController(model, words)
	game, err := c.NewGuess("XXXX", "a")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.AlreadyGuessed.String() {
		t.Fail()
	}
}

func TestGuessBadGuess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	words := mocks.NewMockWords(mockCtrl)

	model.EXPECT().GetGame("XXXX").Return(data.Game{Letters: []string{"t", "e", "s", "t"}, Used: make(map[string]bool), Status: data.BadGuess}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	c := NewProductionController(model, words)
	game, err := c.NewGuess("XXXX", "a")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.BadGuess.String() {
		t.Fail()
	}
}

func TestGuessGoodGuess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	words := mocks.NewMockWords(mockCtrl)

	model.EXPECT().GetGame("XXXX").Return(data.Game{Letters: []string{"t", "e", "s", "t"}, Used: make(map[string]bool), Status: data.BadGuess}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	c := NewProductionController(model, words)
	game, err := c.NewGuess("XXXX", "e")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.GoodGuess.String() {
		t.Fail()
	}
}
