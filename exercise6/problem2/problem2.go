package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	blnc int
	mx   *sync.RWMutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc, &sync.RWMutex{}}
}

func (b *bankAccount) balance() int {
	time.Sleep(readDelay)
	defer b.mx.RUnlock()
	b.mx.RLock()
	return b.blnc
}

func (b *bankAccount) withdraw(amount int) {
	defer b.mx.Unlock()
	b.mx.Lock()
	if b.blnc >= amount {
		b.blnc -= amount
	}
}

func (b *bankAccount) deposit(amount int) {
	defer b.mx.Unlock()
	b.mx.Lock()
	b.blnc += amount
}
