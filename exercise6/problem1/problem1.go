package main

import (
	"fmt"
	"sync"
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

// Balance returns the current balance of the account, ensuring concurrent safety
func (a *bankAccount) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	account := &bankAccount{}

	var wg sync.WaitGroup

	// Concurrently deposit money
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(10)
		}()
	}

	// Concurrently withdraw money
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Withdraw(5)
		}()
	}

	wg.Wait()
	fmt.Println("Final Balance:", account.Balance())
}
