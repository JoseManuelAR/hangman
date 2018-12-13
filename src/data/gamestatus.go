package data

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
	names := [...]string{
		"Initial",
		"GoodGuess",
		"AlreadyGuessed",
		"BadGuess",
		"Lost",
		"Won"}
	if status < Initial || status > Won {
		return "Unknown"
	}
	return names[status]
}
