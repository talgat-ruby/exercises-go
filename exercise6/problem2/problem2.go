package problem2

import (
	"fmt"
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
	time.Sleep(readDelay)
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.blnc
}

func (a *bankAccount) deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.blnc += amount
}

func (a *bankAccount) withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.blnc > amount {
		a.blnc -= amount
		return nil
	}
	return fmt.Errorf("Pupupu, too little money")
}
