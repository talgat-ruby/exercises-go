package problem7

import "fmt"

type BankAccount struct {
	name    string
	balance int
}

type FedexAccount struct {
	name     string
	packages []string
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

type setBalance interface {
	setBalance(balance int)
}

type sendTo interface {
	getName() string
	addPackage(pkg string)
}

func (account *BankAccount) setBalance(amount int) {
	account.balance = account.balance - amount
}

func (account *FedexAccount) getName() string {
	return account.name
}

func (account *FedexAccount) addPackage(pkg string) {
	account.packages = append(account.packages, pkg)
}

func (account *KazPostAccount) getName() string {
	return account.name
}

func (account *KazPostAccount) setBalance(amount int) {
	account.balance = account.balance - amount
}

func (account *KazPostAccount) addPackage(pkg string) {
	account.packages = append(account.packages, pkg)
}

func withdrawMoney(money int, accounts ...setBalance) {
	for _, acc := range accounts {
		acc.setBalance(money)
	}
}

func sendPackagesTo(whom string, accounts ...sendTo) {
	for _, acc := range accounts {
		toWhom := fmt.Sprintf("%s send package to %s", acc.getName(), whom)
		acc.addPackage(toWhom)
	}
}
