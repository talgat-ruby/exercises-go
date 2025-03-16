package main

import (
	"fmt"
	"sync"
	"time"
)

// bankAccount represents a simple bank account with a balance and mutex for concurrent safety
type bankAccount struct {
	balance int
	mu      sync.Mutex
}

// Deposit adds an amount to the account balance, ensuring concurrent safety
func (a *bankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Withdraw subtracts an amount from the account balance if funds are available, ensuring concurrent safety
func (a *bankAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false // insufficient funds
}

// Balance safely returns the current balance of the account
func (a *bankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulate a read delay
	return a.balance
}

func main() {
	account := &bankAccount{}
	account.Deposit(1000)
	fmt.Println("Current Balance:", account.Balance()) // Should print 1000

	account.Withdraw(200)
	fmt.Println("Current Balance after withdrawal:", account.Balance()) // Should print 800
}
