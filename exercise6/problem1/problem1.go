package problem1

import "sync"

type bankAccount struct {
	blnc int
	mu   sync.Mutex
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
	if amount > b.blnc {
		return false
	}
	b.blnc -= amount
	return true
}

func (b *bankAccount) balance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.blnc
}
