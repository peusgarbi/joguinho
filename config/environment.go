package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type environment struct {
	TEXT_SPEED string `validate:"required,numeric"`
}

var Env *environment

func init() {
	godotenv.Load()
	Env = &environment{
		TEXT_SPEED: os.Getenv("TEXT_SPEED"),
	}

	err := validator.New().Struct(Env)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Erro ao carregar variável de ambiente %v -> Erro: '%v', valor recebido: '%v'.", err.StructField(), err.Tag(), err.Value())
			fmt.Println()
		}

		panic("Encerrando aplicação devido ausência de variávies de ambiente!")
	}
}
