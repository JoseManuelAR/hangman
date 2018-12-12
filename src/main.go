package main

import (
	"config"
	"hangman"
	"log"
)

func main() {
	config, error := config.NewConfig()
	if error != nil {
		log.Fatal("Error in config. Exiting program")
	}
	hangman := hangman.NewHangman(config)
	hangman.Start()
}
