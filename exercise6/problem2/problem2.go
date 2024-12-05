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

func (b *bankAccount) balance() int {
	if readDelay > 0 {
		time.Sleep(readDelay)
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.blnc
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
