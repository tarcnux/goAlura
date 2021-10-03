//monitor.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	fmt.Println("O valor da variável comando é: ", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento...")
	//site := "http://www.tarcnux.com.br"
	site := "https://random-status-code.herokuapp.com/" //Retorna http status code aleatório
	resp, _ := http.Get(site)                           //Retorna resp, err, no caso acima igonra o err com o _
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está com problema! HTTP Status code: ", resp.StatusCode)
	}
}
