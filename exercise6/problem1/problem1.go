package problem1

import "sync"

type bankAccount struct {
	blnc int
	mu   sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (a *bankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.blnc += amount
}
func (a *bankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.blnc
}

func (a *bankAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.blnc >= amount {
		a.blnc -= amount
		return true
	}
	return false
}
