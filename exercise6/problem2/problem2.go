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

func (b *bankAccount) balance() int {
	time.Sleep(readDelay)
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.blnc
}

func (b *bankAccount) deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.blnc += amount
}

func (b *bankAccount) withdraw(amount int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.blnc >= amount {
		b.blnc -= amount
		return true
	}
	return false
}
