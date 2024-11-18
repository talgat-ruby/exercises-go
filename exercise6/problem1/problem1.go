package problem1

import "sync"

type bankAccount struct {
	blnc int
	mu   sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (acc *bankAccount) withdraw(money int) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.blnc < money {
		return false
	}
	acc.blnc -= money
	return true
}

func (acc *bankAccount) deposit(money int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.blnc += money
}
