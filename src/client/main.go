package main

import (
	config "config/client"
	"fmt"
)

func main() {
	config := config.NewCliConfig()

	fmt.Println(config.ServerIp())
}
