package data

type GameStatus int

const (
	Initial        string = "Initial"
	GoodGuess      string = "GoodGuess"
	AlreadyGuessed string = "AlreadyGuessed"
	BadGuess       string = "BadGuess"
	Lost           string = "Lost"
	Won            string = "Won"
)

const (
	InitialCode        GameStatus = 0
	GoodGuessCode      GameStatus = 1
	AlreadyGuessedCode GameStatus = 2
	BadGuessCode       GameStatus = 3
	LostCode           GameStatus = 4
	WonCode            GameStatus = 5
)

func (status GameStatus) String() string {
	names := [...]string{
		Initial,
		GoodGuess,
		AlreadyGuessed,
		BadGuess,
		Lost,
		Won}
	if status < InitialCode || status > WonCode {
		return "Unknown"
	}
	return names[status]
}
