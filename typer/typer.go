package typer

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func TypeWriterPrint(text string, cor color.Attribute) {
	c := color.New(cor)
	textSpeed := 20
	keysEvents, _ := keyboard.GetKeys(10)
	defer keyboard.Close()

	color.Yellow("Pressione 'Enter' para pular...")
	for _, char := range text {
		select {
		case key := <-keysEvents:
			if key.Key == keyboard.KeyEnter {
				textSpeed = 0
				c.Print(string(char))
			} else {
				c.Print(string(char))
			}
		default:
			c.Print(string(char))
			time.Sleep(time.Duration(textSpeed) * time.Millisecond)
		}
	}
	fmt.Println()
}

func NormalPrint(text string, cor color.Attribute) {
	c := color.New(cor)
	c.Println(text)
}

func ErrorPrint(text string, err error) {
	color.Red(text)
	yellow := color.New(color.FgYellow)
	yellow.Println(err)
}
