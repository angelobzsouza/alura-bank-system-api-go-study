package main

import "fmt"

type ChekingAccount struct {
	owner string
	agency int
	account int
	balance float64
}

func main() {
	fmt.Println(ChekingAccount{})
}