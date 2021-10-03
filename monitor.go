//monitor.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const repeticoesMonitoramento = 2
const esperarProximoMonitoramento = 2

func main() {

	for {
		exibeIntroducao()
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

	/**
	if comando == 1 {
		fmt.Println("Monitoramento ...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do progama...")
	} else {
		fmt.Println("Não conheço este comando")
	}
	*/
}

func exibeIntroducao() {
	nome, versao := pegaNomeVersao() //Função com retorno de duas variáveis
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func pegaNomeVersao() (string, float64) {
	nome := "Tarcísio"
	versao := 1.1
	return nome, versao
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi: ", comandoLido)
	fmt.Println("") //Pula uma linha
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento...")
	//slice é um tipo de array dinâmico
	sites := lerSitesDoArquivo()

	fmt.Println(sites)

	for i := 0; i <= repeticoesMonitoramento; i++ {
		fmt.Println("") //Pula uma linha
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(esperarProximoMonitoramento * time.Second) //Espera N segundos
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está com problema! HTTP Status code: ", resp.StatusCode)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Erro ao tentar abrir o arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		//fmt.Println(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Erro ao tentar ler o arquivo:", err)
		}
	}
	arquivo.Close()
	return sites
}
