package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	blnc int
	mu   sync.RWMutex // Use RWMutex for better concurrency
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
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

func (b *bankAccount) balance() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	time.Sleep(readDelay)
	return b.blnc
}
