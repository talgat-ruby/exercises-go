package problem2

import (
	"sync"
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	blnc int
	mu   sync.Mutex // мьютекс для защиты баланса
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

// Метод для пополнения баланса
func (a *bankAccount) deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.blnc += amount
}

// Метод для снятия средств с баланса
func (a *bankAccount) withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.blnc >= amount {
		a.blnc -= amount
		return true
	}
	return false
}

// Метод для получения текущего баланса
func (a *bankAccount) balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	time.Sleep(readDelay) // задержка для симуляции времени чтения
	return a.blnc
}
