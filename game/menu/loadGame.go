package menu

import (
	"context"
	"fmt"
	"joguinho/database/sqlite"
	"joguinho/game/utils"
	"joguinho/schema"
	"slices"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

func LoadGame(g *schema.GameContext) error {
	stopLoading := utils.Loading()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	saves, err := sqlite.ListAllSaves(ctx)
	if err != nil {
		return err
	}
	stopLoading()

	choices := parseSavesToLoadStringsOptions(saves)
	var choice string
	err = survey.AskOne(&survey.Select{
		Message: "Qual jogo deseja carregar?",
		Options: choices,
	}, &choice, survey.WithValidator(survey.Required))
	if err != nil {
		return fmt.Errorf("erro ao tentar computar sua escolha do save:\n%v", err)
	}

	save := (*saves)[slices.Index(choices, choice)]

	g.SaveId = save.ID
	g.Nickname = save.Nickname
	g.Currency = save.Currency

	fmt.Println(g)

	g.NextGameFunction = EndOfContent
	return nil
}

func parseSavesToLoadStringsOptions(saves *[]sqlite.Safe) []string {
	var parsedSaves []string
	for _, save := range *saves {
		parsedSaves = append(parsedSaves, fmt.Sprintf("%v- %v", save.ID, save.Nickname))
	}
	return parsedSaves
}
