package problem7

import "fmt"

type BankAccount struct {
	name    string
	balance int
}

func (account *BankAccount) getBalance() int {
	return account.balance
}

func (account *BankAccount) setBalance(amount int) {
	account.balance = account.balance - amount
}

type FedexAccount struct {
	name     string
	packages []string
}

func (account *FedexAccount) getName() string {
	return account.name
}

func (account *FedexAccount) setPackage(pkg string) {
	account.packages = append(account.packages, pkg)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (account *KazPostAccount) getBalance() int {
	return account.balance
}

func (account *KazPostAccount) getName() string {
	return account.name
}

func (account *KazPostAccount) setBalance(amount int) {
	account.balance = account.balance - amount
}

func (account *KazPostAccount) setPackage(pkg string) {
	account.packages = append(account.packages, pkg)
}

type MinusMoney interface {
	getBalance() int
	setBalance(balance int)
}

func withdrawMoney(money int, acc ...MinusMoney) {
	for _, m := range acc {
		m.setBalance(money)
	}
}

type SendSmth interface {
	getName() string
	setPackage(pkg string)
}

func sendPackagesTo(whom string, acc ...SendSmth) {
	for _, m := range acc {
		toWhom := fmt.Sprintf("%s send package to %s", m.getName(), whom)
		m.setPackage(toWhom)
	}
}
