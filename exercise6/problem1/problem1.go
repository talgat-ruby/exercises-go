package problem1

import "sync"

type bankAccount struct {
	blnc int
	m    *sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	m := sync.Mutex{}
	return &bankAccount{blnc, &m}
}

func (b *bankAccount) withdraw(balance int) {
	b.m.Lock()
	if b.blnc > 10 {
		b.blnc = b.blnc - balance
	}
	b.m.Unlock()
}

func (b *bankAccount) deposit(balance int) {
	b.m.Lock()
	b.blnc = b.blnc + balance
	b.m.Unlock()
}
