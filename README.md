# Design
## Applied architecture:
In both programs I have applied an MVC architecture (model view controller view), always considering the SOLID principles.
## Client:
- View: I have developed a simple CLI view (command line interface) that interacts with the main controller that contains the business rules (game actions). If we wanted to implement a more complex interface (for example graphic) the interaction with the controller would be exactly the same through the same interface, it would only change how we present or request the data to the user.
- Controller: I have developed a class with a method for each use case (new game, new guess, list all games, etc). Each action calls a method of a different package per use case. If we wanted to add a new use case (for example delete item) it would be as easy as adding a new package.
- Model: The model in this case is the part that implements communication with the server and obtaining results (remote model). The protocol used in this case has been REST, but new ones could be added, such as GRPC, without affecting existing ones since they have been developed through interfaces. It would simply be
## Server:
- View: In this case the view, that is, where the data enters our system, is a server that receives the commands from a client. As in the client, obviously, the protocol used has been REST. Also in this case we could add more protocols (GRPC for example) without affecting the existing ones.
- Controller: Similar to the client, the controller has the business logic of the server side. Receive the actions to be executed from the view and interact with the model to offer the results to the view again, which will be sent to the client in this case (in other cases it could be presented on the screen)
- Model: I have developed a model in memory (non-persistent) that will save the data of the games in progress. With the same interface a persistent model could be implemented (for example a database)
The server also has an interface that is responsible for generating the random words for each game. In my case, it reads a list of words from a plain text file and selects one at random for each game.

# Implementation

Implmented using go version go1.11 linux/amd64

# Build
- Dependencies: go get ./...
- Building client: go build -o bin/client client/main
- Building server: go build -o bin/server server/main

After these steps, we can run server and client:
- Launch server: bin/server
- Launch client(s): bin/client

We can modify parameters of execution (configuration) by cli parameters (in other cases we can implement the config interface with yaml files, for example)
- -file string
  Words file (default "words/words.txt")
- -ip string
  Listening ip (default "0.0.0.0")
- -model string
  Model type (default "memory")
- -port string
  Listening port (default "8000")
- -remote string
  Remote connection type (default "rest")
- -view string
  View Type (default "cli")
- -words string
  Words Type (default "file")

# Tests & Coverage

go test -v server/controller/newguess -coverprofile=newguess.out

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
ok  	server/controller/newguess	0.003s	coverage: 100.0% of statements

go tool cover -html=newguess.out -o newguess.html

