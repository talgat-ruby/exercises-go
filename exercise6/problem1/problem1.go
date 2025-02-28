package problem1

import "sync"

type bankAccount struct {
	blnc  int
	mutex sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

func (a *bankAccount) deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.blnc += amount
}

func (a *bankAccount) withdraw(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.blnc >= amount {
		a.blnc -= amount
	}
}
