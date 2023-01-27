package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func exibeIntroducao() {
	nome := os.Getenv("USER")
	versao := os.Getenv("VERSION")

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leOpcao() int {
	var opcaoLida int

	fmt.Scan(&opcaoLida)
	fmt.Println("A opção escolhida foi", opcaoLida)

	return opcaoLida
}

func exibeMenu() {
	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func main() {

	exibeIntroducao()
	exibeMenu()

	opcao := leOpcao()

	switch opcao {
	case 1:
		fmt.Println("Monitorando ...")
	case 2:
		fmt.Println("Exibindo logs ...")
	case 3:
		fmt.Println("Saindo do programa ...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida")
		os.Exit(-1)
	}
}
