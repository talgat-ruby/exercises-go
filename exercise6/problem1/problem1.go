package problem1

import "sync"

type bankAccount struct {
	blnc int
	mx   *sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc, &sync.Mutex{}}
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
