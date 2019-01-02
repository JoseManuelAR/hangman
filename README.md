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

# Concurrency issues
Possible concurrency problems could occur in the server with the possibility of multiple clients accessing it simultaneously.
Faced with this possibility two solutions are proposed:
- Goroutines: Each incoming request to the server will be treated in a different goroutine, thus avoiding problems. It is achieved through the mux.Router object and the mapping of functions using HandleFunc
- sync mutex: In the part corresponding to the model, the concurrent access to the data is protected through sync.RWMutex. When accesses are read-only, using RLock / RUnlock and when they are modified or written by Lock / Unlock

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

In this exercise we have only created unit tests for the business logic of 'new guess' on the server, but in a real development it would be advisable (even mandatory) to follow TDD, forcing us to write the tests before the code itself.<br/>
The business logic of the command 'new guess' interacts directly with the interface of the model (to modify the letters guessed in a game), so I had to generate a mock of that interface. This is done by the following line:<br/>

mockgen -source = src / server / model / model.go -destination = src / mocks / mock_model.go -package = mocks<br/>

In this case the unit tests are executed with the following line and we obtain the following output on the screen:<br/>

go test -v server/controller/newguess -coverprofile=newguess.out

=== RUN   TestLetterInWord<br/>
--- PASS: TestLetterInWord (0.00s)<br/>
=== RUN   TestNotLetterInWord<br/>
--- PASS: TestNotLetterInWord (0.00s)<br/>
=== RUN   TestHasWon<br/>
--- PASS: TestHasWon (0.00s)<br/>
=== RUN   TestHasNotWon<br/>
--- PASS: TestHasNotWon (0.00s)<br/>
=== RUN   TestGuessWithEmptyGuess<br/>
--- PASS: TestGuessWithEmptyGuess (0.00s)<br/>
=== RUN   TestGuessToNonExistentGame<br/>
--- PASS: TestGuessToNonExistentGame (0.00s)<br/>
=== RUN   TestGuessToLostGame<br/>
--- PASS: TestGuessToLostGame (0.00s)<br/>
=== RUN   TestGuessToWonGame<br/>
--- PASS: TestGuessToWonGame (0.00s)<br/>
=== RUN   TestGuessAlreadyGuessed<br/>
--- PASS: TestGuessAlreadyGuessed (0.00s)<br/>
=== RUN   TestGuessBadGuess<br/>
--- PASS: TestGuessBadGuess (0.00s)<br/>
=== RUN   TestGuessGoodGuess<br/>
--- PASS: TestGuessGoodGuess (0.00s)<br/>
=== RUN   TestGuessAndLoseGame<br/>
--- PASS: TestGuessAndLoseGame (0.00s)<br/>
=== RUN   TestGuessAndWinGame<br/>
--- PASS: TestGuessAndWinGame (0.00s)<br/>
PASS<br/>
coverage: 100.0% of statements<br/>
ok  	server/controller/newguess	0.003s	coverage: 100.0% of statements<br/>

As you can see I have added a code coverage check of our unit tests, that is, the percentage of our actual code that is exercised by our tests. The output of the coverage can be easily checked and analyzed with the following line:<br/>

go tool cover -html=newguess.out -o newguess.html<br/>

