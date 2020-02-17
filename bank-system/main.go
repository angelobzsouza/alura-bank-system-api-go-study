package main

import (
	"fmt"

	"alura-course-go/bank-system/accounts"
)

func main() {
	silviaAccount := accounts.ChekingAccount{}
	silviaAccount.Owner = "Silvia"
	silviaAccount.Balance = 500

	silviaAccount.Withdraw(400)
	fmt.Println(silviaAccount)
}
