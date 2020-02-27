package accounts

import (
	"../owners"
)

type CheckingAccount struct {
	Owner   owners.Owner
	Agency  int
	Account int
	balance float64
}

func (c *CheckingAccount) Withdraw(withdrawVal float64) bool {
	canWithdraw := withdrawVal <= c.balance && withdrawVal > 0

	if canWithdraw {
		c.balance -= withdrawVal
		return true
	}
	return false
}

func (c *CheckingAccount) Deposit(depositVal float64) bool {
	if depositVal > 0 {
		c.balance += depositVal
		return true
	}
	return false
}

func (c *CheckingAccount) Transfer(transferVal float64, destAccount *CheckingAccount) bool {
	canTransfer := transferVal <= c.balance && transferVal > 0

	if canTransfer {
		destAccount.balance += transferVal
		c.balance -= transferVal
		return true
	}
	return false
}

func (c *CheckingAccount) getBalance() float64 {
	return c.balance
}
