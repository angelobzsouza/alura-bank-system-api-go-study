package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
			showLogs()
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
	registryLog(site, res.StatusCode == 200)
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

func registryLog(site string, online bool) {
	logFile, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var status string
	if online {
		status = "online"
	} else {
		status = "offline"
	}

	logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - Site: " + site + " Status: " + status + "\n")

	logFile.Close()
}

func showLogs() {
	logFile, err := ioutil.ReadFile("log.log")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(logFile))
}
