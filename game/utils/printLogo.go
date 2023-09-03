package utils

import (
	"fmt"
	"os"
)

func PrintLogo() {
	content, _ := os.ReadFile("static/logo.txt")
	fmt.Println(string(content))
}
