package main

import (
	"joguinho/game/menu"
	"joguinho/game/utils"
	"joguinho/schema"
	"time"
)

var (
	gameContext = &schema.GameContext{}
)

func main() {
	stopLoading := utils.Loading()
	time.Sleep(5 * time.Second)
	stopLoading()
	utils.PrintLogo()
	for {
		if gameContext.NextGameFunction == nil {
			break
		}
		err := gameContext.NextGameFunction(gameContext)
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	gameContext.NextGameFunction = menu.MainMenu
}
