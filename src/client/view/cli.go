package view

import (
	"bufio"
	"client/controller"
	"data"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

type cliView struct {
	controller controller.Controller
	writer     *tabwriter.Writer
}

func NewCliView(controller controller.Controller) View {
	return &cliView{controller: controller,
		writer: new(tabwriter.Writer)}
}

func (view cliView) Start(bc chan bool) error {
	view.writer.Init(os.Stdout, 10, 1, 0, ' ', tabwriter.Debug)
	view.mainView()
	bc <- true
	return nil
}

func (view cliView) mainView() {
	for {
		fmt.Println("\n***********************\n" +
			"New Game [n]\n" +
			"Resume games [r]\n" +
			"List games [l]\n" +
			"Exit [x]\n" +
			"***********************\n" +
			"Option:")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)

		if err != nil {
			log.Fatal("Could not read from terminal")
			os.Exit(1)
		}

		switch command {
		case "n":
			view.newGameView()
		case "r":
			view.resumeGameView()
		case "l":
			view.listGamesView()
		case "x":
			os.Exit(0)
		}
	}
}

func (view cliView) newGameView() {
	gameInfo, err := view.controller.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	view.gameView(gameInfo.Id)
}

func (view cliView) resumeGameView() {
	fmt.Println("Enter Id of game to resume:")
	gameId, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	view.gameView(strings.TrimSpace(gameId))
}

func (view cliView) gameView(gameId string) {
	gameInfo, err := view.controller.GetGame(gameId)
	if err != nil {
		log.Fatal(err)
	}
	view.listHeaderView()
	view.listGameView(gameInfo)
	view.writer.Flush()
	for gameInfo.Status != data.Lost && gameInfo.Status != data.Won {
		fmt.Println("Guess a letter for the word:")
		guess, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		if guess == ".x" {
			break
		}
		gameInfo, err = view.controller.NewGuess(gameId, guess)
		if err != nil {
			log.Fatal(err)
		}
		view.listHeaderView()
		view.listGameView(gameInfo)
		view.writer.Flush()
	}
}

func (view cliView) listGamesView() {
	games, err := view.controller.GetGames()
	if err != nil {
		log.Fatal(err)
	}
	view.listHeaderView()
	for _, gameInfo := range games {
		view.listGameView(gameInfo)
	}
	view.writer.Flush()
}

func (view cliView) listHeaderView() {
	fmt.Fprintln(view.writer, "Game Id"+"\t"+"Status"+"\t"+"Word"+"\t")
}

func (view cliView) listGameView(gameInfo data.GameInfo) {
	var status string
	switch {
	case gameInfo.Status == "Initial" || gameInfo.Status == "GoodGuess" || gameInfo.Status == "BadGuess" || gameInfo.Status == "AlreadyGuessed":
		status = "Playing"
	default:
		status = gameInfo.Status
	}
	fmt.Fprintln(view.writer, gameInfo.Id+"\t"+status+"\t"+gameInfo.RevealedWord+"\t")
}
