package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
)

const MST_ADD = "### MST-ADD ###\n"
const MST_NOT_ADD = "### MST-NOT-ADD ###\n"

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
			getStatus()
		case "add":
			addFile()
		default:
			fmt.Println("Comando invalido")
			help()
		}
	} else {
		help()
	}
}

func iniciarRepositorio() {
	fmt.Println()
	_, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Falha ao encontrar path principal: ", err)
		os.Exit(1)
	}
	diretorioAtual := getDiretorioAtual()
	pathMst := path.Join(diretorioAtual, ".mst")
	_, errStatus := os.Stat(pathMst)

	if errStatus != nil {
		fmt.Println("pasta [.mst] não encontrada")
		fmt.Println("Criando repositorio em: ", diretorioAtual)
		os.Mkdir(pathMst, os.ModePerm)
	} else {
		fmt.Println("Repositorio ja existe neste diretorio")
		os.Exit(1)
	}
	fmt.Println("Iniciando repositorio...")
	arquivo, err := criarArquivo(path.Join(pathMst, "diff.mst"))
	if err != nil {
		removerPastaMst()
		fmt.Println("Erro ao criar arquivo diff.mst")
		os.Exit(1)
	}
	// fileBoilerplate, errBoilerplate := getFileBoilerplate()
	// if errBoilerplate != nil {
	// 	removerPastaMst()
	// 	os.Exit(1)
	// }
	// arquivo.Write(fileBoilerplate)
	arquivo.Close()
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
		fmt.Println("Erro ao obter o diretorio atual", err)
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
	diffFile, err := os.ReadFile("./.mst/diff.mst")
	if err != nil {
		fmt.Println("Erro ao ler arquivo diff.mst")
		os.Exit(1)
	}
	diffFileString := string(diffFile)
	// stringsSplit := strings.Split(diffFileString, "\n")
	// fmt.Println(stringsSplit[1])
	dir, _ := os.ReadDir("./")
	var arquivosAdicionados []string
	var arquivosNaoAdicionados []string
	for _, s := range dir {
		if s.IsDir() {
			continue
		}
		hashArquivo, _ := calculateFileHash("./" + s.Name())
		if !strings.Contains(diffFileString, hashArquivo) {
			arquivosNaoAdicionados = append(arquivosNaoAdicionados, color.RedString("	%s", s.Name()))
		} else {
			arquivosAdicionados = append(arquivosAdicionados, color.GreenString("	%s", s.Name()))
		}
	}
	if len(arquivosAdicionados) != 0 {
		fmt.Println("Arquivos adicionados: ")
		for _, s := range arquivosAdicionados {
			fmt.Println(s)
		}

	}
	if len(arquivosNaoAdicionados) != 0 {
		fmt.Println("Arquivos não adicionados: ")
		for _, s := range arquivosNaoAdicionados {
			fmt.Println(s)
		}
	}
}

func calculateFileHash(filePath string) (string, error) {
	// Leitura do conteúdo do arquivo
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	hash.Write(content)
	hashInBytes := hash.Sum(nil)

	hashString := hex.EncodeToString(hashInBytes)
	return hashString, nil
}

func criarArquivo(nomeArquivo string) (*os.File, error) {
	fileCreated, err := os.Create(nomeArquivo)
	if err != nil {
		return nil, err
	}
	return fileCreated, nil
}

func addFile() {
	if len(os.Args) < 3 {
		fmt.Println("Informe o caminho do arquivo")
		os.Exit(1)
	}
	filePath := strings.Trim(os.Args[2], " ")
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Arquivo não encontrado")
		os.Exit(1)
	}
	hashArquivo, err := calculateFileHash(filePath)
	if err != nil {
		fmt.Println("Erro ao calcular hash do arquivo: ", err)
		os.Exit(1)
	}
	fileData, errData := os.ReadFile("./.mst/diff.mst")
	if errData != nil {
		fmt.Println("Erro ao ler arquivo diff.mst")
		os.Exit(1)
	}
	fileDataString := string(fileData)
	if strings.Contains(fileDataString, hashArquivo) {
		fmt.Println("Arquivo ja adicionado")
		os.Exit(1)
	}
	file, _ := os.OpenFile("./.mst/diff.mst", os.O_APPEND|os.O_WRONLY, 0644)
	appendString := append(fileData, []byte("\n")...)
	appendString = append(appendString, []byte(hashArquivo)...)
	file.Write(appendString)
	file.Close()
	fmt.Printf("Arquivo: %s adicionado com sucesso\n", filePath)
}

func removerPastaMst() {
	diretorioAtual := getDiretorioAtual()
	pathMst := path.Join(diretorioAtual, ".mst")
	println("Removendo pasta: ", pathMst)
	os.Remove(pathMst)
}
