package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	monitoramentos = 3
	delay          = 5
)

func exibeIntroducao() {

	nome := os.Getenv("USER")
	versao := os.Getenv("VERSION")

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")

}

func leOpcao() int {

	var opcaoLida int

	fmt.Scan(&opcaoLida)
	//fmt.Println("A opção escolhida foi", opcaoLida)

	return opcaoLida

}

func exibeMenu() {

	fmt.Println("-- Programa de Monitoramento de Sites em Golang --")
	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")
	fmt.Print("Digite uma opção: ")
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func iniciaMonitoramento() {

	fmt.Println("Monitorando ...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("Fim dos testes!")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	//Trata Status Code
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas.")
		registraLogs(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	// Abre o arquivo
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	// Lê o arquivo linha a linha
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	// Fecha o arquivo
	arquivo.Close()
	return sites

}

func registraLogs(site string, status bool) {
	
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	
	arquivo.Close()

}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(arquivo))

}

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		opcao := leOpcao()

		switch opcao {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo logs ...")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Digite novamente.")
			continue
		}
	}
}
