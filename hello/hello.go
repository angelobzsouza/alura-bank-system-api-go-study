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

const monitorQt = 5
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Não conheço esta instrução")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Angelo"
	version := 1.1

	fmt.Println("Olá Sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("0 - Sair")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exebir os logs")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := getSites()

	for i := 0; i < monitorQt; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "carregado corretamente")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code", res.StatusCode)
	}
}

func getSites() []string {
	var sites []string

	file, err := os.Open("./sites.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		site, err := reader.ReadString('\n')
		site = strings.TrimSpace(site)
		if err == io.EOF {
			break
		}
		sites = append(sites, site)
	}

	file.Close()
	return sites
}
