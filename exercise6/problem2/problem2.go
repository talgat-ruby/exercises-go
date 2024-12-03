package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	blnc  int
	mutex sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (b *bankAccount) deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.blnc += amount
}

func (b *bankAccount) withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.blnc >= amount {
		b.blnc -= amount
	}
}

func (b *bankAccount) balance() int {
	time.Sleep(readDelay)
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.blnc
}
