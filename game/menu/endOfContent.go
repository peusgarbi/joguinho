package menu

import (
	"fmt"
	"joguinho/schema"
	"joguinho/typer"

	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func EndOfContent(g *schema.GameContext) error {
	typer.TypeWriterPrint("Infelizmente, por enquanto, o jogo acaba aqui!", color.FgBlue)

	options := []string{"Retornar ao Menu Principal", "Sair do Jogo"}

	var choice string
	err := survey.AskOne(&survey.Select{
		Message: "Escolha:",
		Options: options,
		Default: options[0],
	}, &choice, survey.WithValidator(survey.Required))
	if err != nil {
		return fmt.Errorf("erro em 'EndOfContent': falha ao processar sua escolha:\n%v", err)
	}
	choiceIndex := slices.Index(options, choice)

	if choiceIndex == 0 {
		g.NextGameFunction = MainMenu
	} else if choiceIndex == 1 {
		g.NextGameFunction = CloseGame
	} else {
		return fmt.Errorf("erro em 'EndOfContent': você conseguiu selecionar uma opção inválida, parabéns")
	}

	return nil
}
