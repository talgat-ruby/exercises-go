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

func (acc *bankAccount) deposit(amount int) {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	acc.blnc += amount
}

func (acc *bankAccount) withdraw(amount int) {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	if acc.blnc >= amount {
		acc.blnc -= amount
	}
}

func (acc *bankAccount) balance() int {
	time.Sleep(readDelay)
	acc.mutex.Lock()
	defer acc.mutex.Unlock()

	return acc.blnc
}
