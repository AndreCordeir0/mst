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
		fmt.Println("pasta [.mst] nao encontrada")
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
	for _, s := range dir {
		if s.IsDir() {
			continue
		}
		hashArquivo, _ := calculateFileHash("./" + s.Name())
		if !strings.Contains(diffFileString, hashArquivo) {
			fmt.Println(color.RedString("Arquivo nao adicionado: %s", s.Name()))
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

func getFileBoilerplate() ([]byte, error) {
	byteText := []byte(MST_ADD + MST_NOT_ADD)
	dir, _ := os.ReadDir("./")
	for _, d := range dir {
		if d.IsDir() {
			continue
		}
		hashArquivo, err := calculateFileHash("./" + d.Name())
		if err != nil {
			fmt.Println("Erro ao calcular hash do arquivo: ", err)
			return nil, err
		}
		fmt.Println("Nome do arquivo: %s Hash do arquivo %s ", d.Name(), hashArquivo)
		byteText = append(byteText, []byte(hashArquivo)...)
		byteText = append(byteText, []byte("\n")...)
		fmt.Println("\n")
	}
	fmt.Println(dir)
	return byteText, nil
}

func removerPastaMst() {
	diretorioAtual := getDiretorioAtual()
	pathMst := path.Join(diretorioAtual, ".mst")
	println("Removendo pasta: ", pathMst)
	os.Remove(pathMst)
}
