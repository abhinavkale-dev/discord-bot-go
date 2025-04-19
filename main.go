package main

import (
	"fmt"

	"github.com/abhinavkale-dev/go-discord-bot/bot"
	"github.com/abhinavkale-dev/go-discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
