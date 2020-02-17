package accounts

type ChekingAccount struct {
	Owner   string
	Agency  int
	Account int
	Balance float64
}

func (c *ChekingAccount) Withdraw(withdrawVal float64) bool {
	canWithdraw := withdrawVal <= c.Balance && withdrawVal > 0

	if canWithdraw {
		c.Balance -= withdrawVal
		return true
	}
	return false
}

func (c *ChekingAccount) Deposit(depositVal float64) bool {
	if depositVal > 0 {
		c.Balance += depositVal
		return true
	}
	return false
}

func (c *ChekingAccount) Transfer(transferVal float64, destAccount *ChekingAccount) bool {
	canTransfer := transferVal <= c.Balance && transferVal > 0

	if canTransfer {
		destAccount.Balance += transferVal
		c.Balance -= transferVal
		return true
	}
	return false
}
