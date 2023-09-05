package menu

import (
	"context"
	"joguinho/database/sqlite"
	"joguinho/game/utils"
	"joguinho/schema"
	"log"
)

func Initialization(g *schema.GameContext) error {
	stopLoading := utils.Loading()
	ctx := context.Background()
	err := sqlite.LocalSqlite.Migrate(ctx)
	if err != nil {
		log.Fatalf("fodeu 1: %v", err)
	}
	stopLoading()
	utils.PrintLogo()
	g.NextGameFunction = MainMenu
	return nil
}
