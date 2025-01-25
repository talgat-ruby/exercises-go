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
	b.mu.Lock()
	defer b.mu.Unlock()
	// time.Sleep(readDelay)
	return b.blnc
}
func (b *bankAccount) deposit(plus int) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.blnc += plus

}
func (b *bankAccount) withdraw(minus int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.blnc >= minus {
		b.blnc -= minus
		return true
	}

	return false
}
