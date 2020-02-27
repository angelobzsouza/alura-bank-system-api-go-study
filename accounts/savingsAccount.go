package accounts

import "../owners"

type SavingsAccount struct {
	Owner     owners.Owner
	Agency    int
	Account   int
	Operation int
	balance   float64
}

func (s *SavingsAccount) Withdraw(withdrawVal float64) bool {
	canWithdraw := withdrawVal <= s.balance && withdrawVal > 0

	if canWithdraw {
		s.balance -= withdrawVal
		return true
	}
	return false
}

func (s *SavingsAccount) Deposit(depositVal float64) bool {
	if depositVal > 0 {
		s.balance += depositVal
		return true
	}
	return false
}
