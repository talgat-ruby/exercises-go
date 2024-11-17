package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	sync.Mutex
	blnc int
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc:blnc}
}

func (b *bankAccount) balance() int {
	time.Sleep(readDelay)
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.blnc
}

func (b *bankAccount) deposit(cash int) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.blnc += cash
}

func (b *bankAccount) withdraw(cash int) bool {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	if b.blnc < cash {
		return false
	}
	b.blnc -= cash
	return true
}

func (b *bankAccount) getBalance() int {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	return b.blnc
}
