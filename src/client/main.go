package main

import (
	config "config/client"
	"fmt"

	"github.com/go-resty/resty"
)

func main() {
	config := config.NewCliConfig()
	fmt.Println(config.ServerIp())

	resp, err := resty.R().SetHeader("Content-Type", "application/json").
		Post("http://" + config.ServerIp() + ":8000/games")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
