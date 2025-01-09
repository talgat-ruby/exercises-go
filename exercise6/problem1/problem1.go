package problem1

import "sync"

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
