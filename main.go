package main

import (
	"fmt"

	"./accounts"
)

func PayBill(account verifyAccount, billValue float64) {
	account.Withdraw(billValue)
}

type verifyAccount interface {
	Withdraw(withdrawVal float64) bool
}

func main() {
	dennisAccount := accounts.SavingsAccount{}
	jayAccount := accounts.CheckingAccount{}
	dennisAccount.Deposit(500)
	jayAccount.Deposit(1000)
	PayBill(&dennisAccount, 400)
	PayBill(&jayAccount, 800)

	fmt.Println(dennisAccount)
	fmt.Println(jayAccount)
}
