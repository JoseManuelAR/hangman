# hangman

go version go1.11 linux/amd64

# tests & coverage

go test -v newguess -coverprofile=newguess.out

=== RUN   TestLetterInWord
--- PASS: TestLetterInWord (0.00s)
=== RUN   TestNotLetterInWord
--- PASS: TestNotLetterInWord (0.00s)
=== RUN   TestHasWon
--- PASS: TestHasWon (0.00s)
=== RUN   TestHasNotWon
--- PASS: TestHasNotWon (0.00s)
=== RUN   TestGuessWithEmptyGuess
--- PASS: TestGuessWithEmptyGuess (0.00s)
=== RUN   TestGuessToNonExistentGame
--- PASS: TestGuessToNonExistentGame (0.00s)
=== RUN   TestGuessToLostGame
--- PASS: TestGuessToLostGame (0.00s)
=== RUN   TestGuessToWonGame
--- PASS: TestGuessToWonGame (0.00s)
=== RUN   TestGuessAlreadyGuessed
--- PASS: TestGuessAlreadyGuessed (0.00s)
=== RUN   TestGuessBadGuess
--- PASS: TestGuessBadGuess (0.00s)
=== RUN   TestGuessGoodGuess
--- PASS: TestGuessGoodGuess (0.00s)
=== RUN   TestGuessAndLoseGame
--- PASS: TestGuessAndLoseGame (0.00s)
=== RUN   TestGuessAndWinGame
--- PASS: TestGuessAndWinGame (0.00s)
PASS
coverage: 100.0% of statements
ok  	newguess	0.003s	coverage: 100.0% of statements

go tool cover -html=newguess.out -o newguess.html

# build

go build server
go build client

