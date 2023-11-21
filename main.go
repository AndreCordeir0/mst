package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	argumentos := os.Args
	if len(argumentos) > 1 {
		comando := argumentos[1]
		switch comando {
		case "help":
			help()
		case "list":
			help()
		case "status":
			getDiretorioAtual()
		default:
			fmt.Println("Comando invalido")
			help()
		}
	} else {
		help()
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(color.RedString(os.Args[i]))
	}
}

func help() {
	fmt.Println("Comandos: mst [help] [list] [status] [init]")
	fmt.Println("init - Para iniciar um novo repositorio no diretorio local")
	fmt.Println("list - Lista os arquivos e pastas do diretorio atual")
	fmt.Println("status - Mostra o status do repositorio atual")
}

func getDiretorioAtual() string {
	diretorio, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretorio atual")
		os.Exit(1)
	} else {
		fmt.Println("Diretorio atual: ", diretorio)
	}
	return diretorio
}
