package problem1

import (
	"sync"
)

type bankAccount struct {
	mu   sync.Mutex
	blnc int
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (a *bankAccount) deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.blnc += amount
}

func (a *bankAccount) withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.blnc >= amount {
		a.blnc -= amount
	}
}
