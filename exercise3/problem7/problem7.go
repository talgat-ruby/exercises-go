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

func withdrawMoney(amount int, accounts ...interface{}) {
	for _, account := range accounts {
		if a, ok := account.(*BankAccount); ok {
			if a.balance >= amount {
				a.balance -= amount
			}
		} else if a, ok := account.(*KazPostAccount); ok {
			if a.balance >= amount {
				a.balance -= amount
			}
		}
	}
}

func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		if a, ok := account.(*FedexAccount); ok {
			message := fmt.Sprintf("%s send package to %s", a.name, recipient)
			a.packages = append(a.packages, message)
		} else if a, ok := account.(*KazPostAccount); ok {
			message := fmt.Sprintf("%s send package to %s", a.name, recipient)
			a.packages = append(a.packages, message)
		}
	}
}
