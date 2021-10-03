//monitor.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const repeticoesMonitoramento = 3
const esperarProximoMonitoramento = 5

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
	sites := []string{"http://www.tarcnux.com.br",
		"https://random-status-code.herokuapp.com/",
		"https://mestreemqueijos.com.br",
		"https://ideiavegana.com.br"}

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
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está com problema! HTTP Status code: ", resp.StatusCode)
	}
}
