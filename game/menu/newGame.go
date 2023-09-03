package menu

import (
	"errors"
	"fmt"
	"joguinho/schema"
	"joguinho/typer"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

type reason string

const (
	pobreFodido reason = "Você era um pobre fudido e decidiu melhorar sua vida."
	deusVult    reason = "Você é o escudo e a espada da igreja e acredita na instituição. Deus Vult!"
	cetico      reason = "Você acredita em Deus, mas perdeu a fé na instituição. Ainda assim quer dedicar sua vida à missão de Cristo."
	pauMandado  reason = "Sua família te obrigou a virar um Seminarista."
	inquisitor  reason = "Você sempre teve uma natureza Inquisitiva sobre as coisas do mundo."
	ganancioso  reason = "Você só se importa com o Poder e o Status que a instituição pode lhe oferecer."
)

func returnAllReasonsAsStrings() []string {
	return []string{
		string(pobreFodido),
		string(deusVult),
		string(cetico),
		string(pauMandado),
		string(inquisitor),
		string(ganancioso),
	}
}

func parseReasonStringToEnum(stringReason string) reason {
	switch stringReason {
	case string(pobreFodido):
		return pobreFodido
	case string(deusVult):
		return deusVult
	case string(cetico):
		return cetico
	case string(pauMandado):
		return pauMandado
	case string(inquisitor):
		return inquisitor
	case string(ganancioso):
		return ganancioso
	default:
		return pobreFodido
	}
}

func NewGame(g *schema.GameContext) error {
	err := survey.AskOne(
		&survey.Input{
			Message: "Qual o seu nome?",
			Help:    "Esse será seu nome durante toda a jornada",
		}, &g.Nickname,
		survey.WithValidator(func(ans interface{}) error {
			if ans.(string) == "" {
				//lint:ignore ST1005 this error message should render as capitalized
				return errors.New("Não é possível prosseguir sem nome!")
			}
			return nil
		}),
	)
	if err != nil {
		return fmt.Errorf("erro ao tentar pegar seu nome:\n%v", err)
	}

	fmt.Println(color.CyanString("Olá,"), color.GreenString(g.Nickname), color.CyanString(", tudo bem com você?"))
	color.Cyan("O jogo ainda está em desenvolvimento, mas você pode contribuir!")

	var isSure bool
	err = survey.AskOne(&survey.Confirm{
		Message: "Deseja prosseguir?",
	}, &isSure)
	if err != nil {
		return fmt.Errorf("erro ao tentar pegar sua confirmação:\n%v", err)
	}
	if !isSure {
		typer.NormalPrint("Jogo abortado, retornando ao menu principal...", color.FgRed)
		g.NextGameFunction = MainMenu
		return nil
	}

	typer.TypeWriterPrint("Você é um jovem seminarista prestes a se formar Padre da Igreja Católica Apostólica Romana no auge do século XIII. Ao olhar para o seu passado, questões vem a sua cabeça se você tomou as melhores decisões...", color.FgMagenta)

	var questions = []*survey.Question{
		{
			Name: "reason",
			Prompt: &survey.Select{
				Message: "Por que você decidiu entrar na Igreja?",
				Options: returnAllReasonsAsStrings(),
			},
			Validate: survey.Required,
		},
		{
			Name: "study",
			Prompt: &survey.Select{
				Message: "Como foram seus primeiros anos de estudo?",
				Options: []string{
					"Você se dedicou muito ao estudo dos Demônios.",
					"Você se dedicou ao estudo da Liturgia.",
				},
			},
			Validate: survey.Required,
			Transform: func(ans interface{}) (newAns interface{}) {
				println(ans)
				return ans
			},
		},
	}
	answers := struct {
		Reason string
		Study  string
	}{}
	err = survey.Ask(questions, &answers)
	if err != nil {
		return fmt.Errorf("erro ao tentar pegar suas respostas:\n%v", err)
	}

	g.NextGameFunction = MainMenu
	return nil
}
