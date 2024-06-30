package problem7

import (
	"fmt"
)

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

func withdrawMoney(amount int, fromAccount BankAccount, toAccount *KazPostAccount) {
	fromAccount.balance -= amount
	toAccount.balance += amount
}

func sendPackagesTo(receiver string, fromAccount FedexAccount, toAccount *KazPostAccount) {
	packageMessage := fmt.Sprintf("%s send package to %s", fromAccount.name, receiver)
	fromAccount.packages = append(fromAccount.packages, packageMessage)
	toAccount.packages = append(toAccount.packages, packageMessage)
