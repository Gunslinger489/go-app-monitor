package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"time"
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

func iniciaMonitoramento() {

	fmt.Println("Monitorando ...")
	sites := []string{"https://butia.rs.gov.br", "https://google.com.br", "https://terra.com.br", "https://httpstat.us/404"}

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(5 * time.Second)
	}
	
	fmt.Println("")

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas.")
	}
}

func main() {

	exibeIntroducao()
	exibeMenu()

	opcao := leOpcao()

	switch opcao {
	case 1:
		iniciaMonitoramento()
	case 2:
		fmt.Println("Exibindo logs ...")
	case 3:
		fmt.Println("Saindo do programa ...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida. Digite novamente.")
		os.Exit(-1)
	}
}
