package problem1

import "sync"

type bankAccount struct {
	mu   sync.Mutex
	blnc int
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
