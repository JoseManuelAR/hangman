package data

const (
	GamePlaying string = "Playing"
	GameLost    string = "Lost"
	GameWon     string = "Won"
)

type GameStatus int

const (
	Initial        GameStatus = 0
	GoodGuess      GameStatus = 1
	AlreadyGuessed GameStatus = 2
	BadGuess       GameStatus = 3
	Lost           GameStatus = 4
	Won            GameStatus = 5
)

func (status GameStatus) String() string {
	switch status {
	case Lost:
		return GameLost
	case Won:
		return GameWon
	default:
		return GamePlaying
	}
}
