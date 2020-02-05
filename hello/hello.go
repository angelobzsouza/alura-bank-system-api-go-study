package main

import (
	"fmt"
	"os"
)

func main() {
	showIntro()
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
}
