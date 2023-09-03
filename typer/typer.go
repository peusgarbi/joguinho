package typer

import (
	"fmt"
	"joguinho/config"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func TypeWriterPrint(text string, cor color.Attribute) {
	c := color.New(cor)
	textSpeed, _ := strconv.Atoi(config.Env.TEXT_SPEED)
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
