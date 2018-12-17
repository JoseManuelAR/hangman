package main

import (
	"config"
	"server/hangman"
)

func main() {
	config := config.NewCliConfig()
	hangman := hangman.NewHangman(config)
	hangman.Start()
}
