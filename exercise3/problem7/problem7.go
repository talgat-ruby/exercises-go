package main

import "fmt"

// Withdrawable interface allows withdrawing money
type Withdrawable interface {
	Withdraw(amount int) bool
}

// Sendable interface allows sending packages
type Sendable interface {
	SendPackage(to string)
}

// BankAccount represents a bank account with a balance
type BankAccount struct {
	name    string
	balance int
}

// FedexAccount represents a FedEx account that can send packages
type FedexAccount struct {
	name     string
	packages []string
}

// KazPostAccount represents a postal account with balance and package capabilities
type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

// Withdraw method for BankAccount
func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("%s withdrew %d. Remaining balance: %d\n", b.name, amount, b.balance)
		return true
	}
	fmt.Printf("%s has insufficient funds to withdraw %d.\n", b.name, amount)
	return false
}

// Withdraw method for KazPostAccount
func (k *KazPostAccount) Withdraw(amount int) bool {
	if k.balance >= amount {
		k.balance -= amount
		fmt.Printf("%s withdrew %d. Remaining balance: %d\n", k.name, amount, k.balance)
		return true
	}
	fmt.Printf("%s has insufficient funds to withdraw %d.\n", k.name, amount)
	return false
}

// SendPackage method for FedexAccount
func (f *FedexAccount) SendPackage(to string) {
	f.packages = append(f.packages, to)
	fmt.Printf("%s sent a package to %s via FedEx.\n", f.name, to)
}

// SendPackage method for KazPostAccount
func (k *KazPostAccount) SendPackage(to string) {
	k.packages = append(k.packages, to)
	fmt.Printf("%s sent a package to %s via KazPost.\n", k.name, to)
}

// withdrawMoney accepts any account that supports withdrawal and attempts to withdraw the specified amount
func withdrawMoney(amount int, accounts ...Withdrawable) {
	for _, account := range accounts {
		account.Withdraw(amount)
	}
}

// sendPackagesTo accepts any account that can send packages and sends a package to the specified recipient
func sendPackagesTo(to string, accounts ...Sendable) {
	for _, account := range accounts {
		account.SendPackage(to)
	}
}

func main() {
	// Create accounts
	normanOsborne := &BankAccount{
		name:    "Norman Osborne",
		balance: 1_000_000,
	}
	peterParker := &FedexAccount{
		name:     "Peter Parker",
		packages: []string{},
	}
	auntMay := &KazPostAccount{
		name:     "Aunt May",
		balance:  200,
		packages: []string{},
	}

	// Test withdrawal
	withdrawMoney(10, normanOsborne, auntMay) // Multiple accounts can withdraw

	// Test sending packages
	sendPackagesTo("Mary Jane", peterParker, auntMay) // Multiple accounts can send packages
}
