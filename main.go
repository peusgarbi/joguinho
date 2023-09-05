package main

import (
	"joguinho/game/menu"
	"joguinho/game/utils"
	"joguinho/schema"
	"joguinho/typer"

	"github.com/inancgumus/screen"
)

var (
	gameContext = &schema.GameContext{}
)

func main() {
	for {
		err := utils.SaveGame(gameContext)
		if err != nil {
			typer.ErrorPrint("Erro ao tentar salvar seu jogo:", err)
			break
		}
		if gameContext.NextGameFunction == nil {
			break
		}
		err = gameContext.CallNextGameFunction()
		if err != nil {
			typer.ErrorPrint("Erro ao tentar chamar a próxima parte do jogo:", err)
			break
		}
	}
}

func init() {
	screen.Clear()
	screen.MoveTopLeft()
	gameContext.NextGameFunction = menu.Initialization
}
