package main

import (
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
)

func main() {
	argumentos := os.Args
	if len(argumentos) > 1 {
		comando := argumentos[1]
		switch comando {
		case "help":
			help()
		case "init":
			iniciarRepositorio()
		case "list":
			listarRepositorios()
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

func iniciarRepositorio() {
	fmt.Println()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Falha ao encontrar path principal: ", err)
		os.Exit(1)
	}
	path := path.Join(homeDir, ".mst")
	_, errStatus := os.Lstat(path)
	if errStatus != nil {
		fmt.Println("pasta [.mst] nao encontrada")
		fmt.Println("Criando repositorio em: ", path)
		os.Mkdir(path, os.ModePerm)
	}

	//
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

func listarRepositorios() {
	fmt.Println("Nenhum repositorio encontrado...")
}

func getStatus() {

}
