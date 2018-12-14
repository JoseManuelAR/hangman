package main

import (
	config "config/server"
	"hangman"
)

func main() {
	config := config.NewCliConfig()
	hangman := hangman.NewHangman(config)
	hangman.Run()
}
