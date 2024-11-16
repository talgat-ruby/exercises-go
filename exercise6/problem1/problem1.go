package problem1

import "sync"

type bankAccount struct {
	blnc int
	mu   sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc: blnc}
}

// Метод для пополнения баланса
func (a *bankAccount) deposit(amount int) {
	a.mu.Lock()         // блокировка мьютекса перед изменением баланса
	defer a.mu.Unlock() // разблокировка мьютекса после операции
	a.blnc += amount
}

// Метод для снятия средств с баланса
func (a *bankAccount) withdraw(amount int) bool {
	a.mu.Lock()         // блокировка мьютекса перед изменением баланса
	defer a.mu.Unlock() // разблокировка мьютекса после операции

	// проверка, достаточно ли средств для снятия
	if a.blnc >= amount {
		a.blnc -= amount
		return true
	}
	return false
}
