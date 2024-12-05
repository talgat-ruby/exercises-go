package problem1

import (
	"sync"
)

type bankAccount struct {
	blnc int
	mtx  *sync.Mutex
}

func (b *bankAccount) withdraw(s int) {
	defer b.mtx.Unlock()
	b.mtx.Lock()
	if s > b.blnc {
		return
	}
	b.blnc -= s

}

func (b *bankAccount) deposit(s int) {
	defer b.mtx.Unlock()
	b.mtx.Lock()
	b.blnc += s
}

func newAccount(blnc int) *bankAccount {
	var m sync.Mutex
	return &bankAccount{blnc, &m}
}
