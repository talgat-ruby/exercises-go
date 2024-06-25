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

func withdrawMoney(amount int, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *BankAccount:
			acc.balance -= amount
		case *KazPostAccount:
			acc.balance -= amount
		}
	}
}

func sendPackagesTo(receiver string, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *FedexAccount:
			acc.packages = append(acc.packages, fmt.Sprintf("%s send package to %s", acc.name, receiver))
			fmt.Printf("Sent package from FedexAccount %s to %s\n", acc.name, receiver)
		case *KazPostAccount:
			acc.packages = append(acc.packages, fmt.Sprintf("%s send package to %s", acc.name, receiver))
			fmt.Printf("Sent package from KazPostAccount %s to %s\n", acc.name, receiver)
		}
	}
}
