package main

import (
	"joguinho/game/menu"
	"joguinho/game/utils"
	"joguinho/schema"
)

var (
	gameContext = &schema.GameContext{}
)

func main() {
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
