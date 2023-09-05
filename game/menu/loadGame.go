package menu

import (
	"context"
	"fmt"
	"joguinho/database/sqlite"
	"joguinho/game/utils"
	"joguinho/schema"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

func LoadGame(g *schema.GameContext) error {
	stopLoading := utils.Loading()

	db, err := sqlite.LocalSqlite.Connect()
	if err != nil {
		return err
	}
	queries := sqlite.New(db)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	saves, err := queries.GetAllSaves(ctx)
	if err != nil {
		return err
	}
	stopLoading()

	var save string
	err = survey.AskOne(&survey.Select{
		Message: "Qual jogo deseja carregar?",
		Options: parseSavesToLoadStringsOptions(saves),
	}, &save, survey.WithValidator(survey.Required))
	if err != nil {
		return fmt.Errorf("erro ao tentar computar sua escolha do save:\n%v", err)
	}

	fmt.Println(save)
	g.NextGameFunction = MainMenu
	return nil
}

func parseSavesToLoadStringsOptions(saves []sqlite.Safe) []string {
	var parsedSaves []string
	for _, save := range saves {
		parsedSaves = append(parsedSaves, fmt.Sprintf("%v- %v", save.ID, save.Nickname))
	}
	return parsedSaves
}
