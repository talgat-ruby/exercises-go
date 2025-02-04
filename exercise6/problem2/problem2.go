package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	blnc int
	mu   sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (acc *bankAccount) balance() int {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.blnc
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
