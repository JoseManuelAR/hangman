package main

import (
	"client/hangman"
	"config"
)

func main() {
	config := config.NewCliConfig()
	hangman := hangman.NewHangman(config)
	hangman.Start()

	// fmt.Println(config.Ip())

	// resp, err := resty.R().SetHeader("Content-Type", "application/json").
	// 	Post("http://" + config.Ip() + ":" + config.Port() + "/hangman/v1/games")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// gameInfo := data.GameInfo{}
	// json.Unmarshal([]byte(resp.String()), &gameInfo)
	// fmt.Println(gameInfo.Id)
}
