package accounts

import (
	"../owners"
)

type ChekingAccount struct {
	Owner   owners.Owner
	Agency  int
	Account int
	balance float64
}

func (c *ChekingAccount) Withdraw(withdrawVal float64) bool {
	canWithdraw := withdrawVal <= c.balance && withdrawVal > 0

	if canWithdraw {
		c.balance -= withdrawVal
		return true
	}
	return false
}

func (c *ChekingAccount) Deposit(depositVal float64) bool {
	if depositVal > 0 {
		c.balance += depositVal
		return true
	}
	return false
}

func (c *ChekingAccount) Transfer(transferVal float64, destAccount *ChekingAccount) bool {
	canTransfer := transferVal <= c.balance && transferVal > 0

	if canTransfer {
		destAccount.balance += transferVal
		c.balance -= transferVal
		return true
	}
	return false
}

func (c *ChekingAccount) getBalance() float64 {
	return c.balance
}
