package problem7

import "fmt"

type BankAccount struct {
	name    string
	balance float64
}

type FedexAccount struct {
	name     string
	packages []string
}

type KazPostAccount struct {
	name     string
	balance  float64
	packages []string
}

func withdrawMoney(amount float64, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *BankAccount:
			if acc.balance >= amount {
				acc.balance -= amount
				fmt.Printf("Withdrew %.2f from %s's bank account. New balance: %.2f\n", amount, acc.name, acc.balance)
			} else {
				fmt.Printf("Insufficient funds in %s's bank account to withdraw %.2f\n", acc.name, amount)
			}
		case *KazPostAccount:
			if acc.balance >= amount {
				acc.balance -= amount
				fmt.Printf("Withdrew %.2f from %s's KazPost account. New balance: %.2f\n", amount, acc.name, acc.balance)
			} else {
				fmt.Printf("Insufficient funds in %s's KazPost account to withdraw %.2f\n", acc.name, amount)
			}
		}
	}
}
func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *FedexAccount:
			acc.packages = append(acc.packages, recipient)
			fmt.Printf("Sent a package to %s from %s's FedEx account. Total packages sent: %v\n", recipient, acc.name, acc.packages)

		case *KazPostAccount:
			acc.packages = append(acc.packages, recipient)
			fmt.Printf("Sent a package to %s from %s's KazPost account. Total packages sent: %v\n", recipient, acc.name, acc.packages)
		}
	}
}
