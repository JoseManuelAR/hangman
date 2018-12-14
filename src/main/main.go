package main

import (
	"config"
	"hangman"
)

func main() {
	config := config.NewCliConfig()
	hangman := hangman.NewHangman(config)
	hangman.Run()
}
