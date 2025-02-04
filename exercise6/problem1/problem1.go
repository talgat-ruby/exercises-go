package problem1

import "sync"

type bankAccount struct {
	blnc int
	mu   sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (acc *bankAccount) deposit(amount int) {
	if amount > 0 {
		acc.mu.Lock()
		defer acc.mu.Unlock()
		acc.blnc += amount
	}
}

func (acc *bankAccount) withdraw(amount int) {

	if amount <= acc.blnc {
		acc.mu.Lock()
		defer acc.mu.Unlock()
		acc.blnc -= amount
	}
}
