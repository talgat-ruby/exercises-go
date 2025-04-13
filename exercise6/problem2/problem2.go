package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

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

func (a *bankAccount) balance() int {
	time.Sleep(readDelay)
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.blnc
}
