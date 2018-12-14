package newguess

import (
	"data"
	"errs"
	gomock "github.com/golang/mock/gomock"
	"mocks"
	"testing"
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

func TestHasWon(t *testing.T) {
	letters := []string{"t", "e", "s", "t"}
	used := make(map[string]bool)
	used["t"] = true
	used["e"] = true
	used["s"] = true
	won := hasWon(letters, used)
	if !won {
		t.Errorf("Game should be won")
	}
}

func TestHasNotWon(t *testing.T) {
	letters := []string{"t", "e", "s", "t"}
	used := make(map[string]bool)
	used["t"] = true
	used["s"] = true
	won := hasWon(letters, used)
	if won {
		t.Errorf("Game should not be won")
	}
}

func TestGuessToNonExistentGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	model.EXPECT().GetGame("XXXX").Return(data.Game{}, errs.ErrGameNotFound)
	_, err := NewGuess(model, "XXXX", "a")

	if err != errs.ErrGameNotFound {
		t.Fail()
	}
}

func TestGuessToLostGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	model.EXPECT().GetGame("XXXX").Return(data.Game{Status: data.Lost}, nil)
	game, err := NewGuess(model, "XXXX", "a")

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
	model.EXPECT().GetGame("XXXX").Return(data.Game{Status: data.Won}, nil)
	game, err := NewGuess(model, "XXXX", "a")

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
	used := make(map[string]bool)
	used["a"] = true
	model.EXPECT().GetGame("XXXX").Return(data.Game{Status: data.BadGuess, Used: used}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	game, err := NewGuess(model, "XXXX", "a")

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
	model.EXPECT().GetGame("XXXX").Return(data.Game{Letters: []string{"t", "e", "s", "t"}, Used: make(map[string]bool), Status: data.BadGuess}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	game, err := NewGuess(model, "XXXX", "a")

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
	model.EXPECT().GetGame("XXXX").Return(data.Game{Letters: []string{"t", "e", "s", "t"}, Used: make(map[string]bool), Status: data.BadGuess}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	game, err := NewGuess(model, "XXXX", "e")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.GoodGuess.String() {
		t.Fail()
	}
}

func TestGuessAndLoseGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	used := make(map[string]bool)
	used["t"] = true
	used["s"] = true
	model.EXPECT().GetGame("XXXX").Return(data.Game{Letters: []string{"t", "e", "s", "t"}, Used: used, Status: data.BadGuess, TurnsLeft: 1}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	game, err := NewGuess(model, "XXXX", "x")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.Lost.String() {
		t.Fail()
	}
}

func TestGuessAndWinGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockModel(mockCtrl)
	used := make(map[string]bool)
	used["t"] = true
	used["s"] = true
	model.EXPECT().GetGame("XXXX").Return(data.Game{Letters: []string{"t", "e", "s", "t"}, Used: used, Status: data.BadGuess, TurnsLeft: 1}, nil)
	model.EXPECT().UpdateGame("XXXX", gomock.Any()).Return(nil)
	game, err := NewGuess(model, "XXXX", "e")

	if err != nil {
		t.Fail()
	}
	if game.Status != data.Won.String() {
		t.Fail()
	}
}