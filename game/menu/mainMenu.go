package menu

import (
	"fmt"
	"joguinho/schema"
	"joguinho/typer"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

type mainMenuChoice string

const (
	novoJogo     mainMenuChoice = "Novo Jogo"
	carregarJogo mainMenuChoice = "Carregar Jogo"
	sair         mainMenuChoice = "Sair"
)

func returnAllMainMenuChoisesAsStrings() []string {
	return []string{
		string(novoJogo),
		string(carregarJogo),
		string(sair),
	}
}

func MainMenu(g *schema.GameContext) error {
	var choice string
	prompt := &survey.Select{
		Message: "Menu Principal",
		Options: returnAllMainMenuChoisesAsStrings(),
		Default: string(novoJogo),
	}
	err := survey.AskOne(prompt, &choice, survey.WithValidator(survey.Required))
	if err != nil {
		log.Fatalf("Erro ao tentar computar sua escolha do menu principal:\n%v", err)
	}

	switch choice {
	case string(novoJogo):
		fmt.Println()
		g.NextGameFunction = NewGame
		return nil
	case string(carregarJogo):
		fmt.Println()
		g.NextGameFunction = LoadGame
		return nil
	case string(sair):
		fmt.Println()
		g.NextGameFunction = CloseGame
		return nil
	default:
		return fmt.Errorf("você conseguiu escolher uma opção inexistente, parabéns")
	}
}

func CloseGame(g *schema.GameContext) error {
	typer.NormalPrint("Fechando Joguinho...", color.FgHiRed)
	g.NextGameFunction = nil
	return nil
}
