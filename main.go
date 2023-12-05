package main

import (
	"app-linha-comando/app"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Aplicação linha de comando")
	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
}
