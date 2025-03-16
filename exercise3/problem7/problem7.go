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
		switch v := account.(type) {
		case *BankAccount:
			if v.balance >= amount {
				v.balance -= amount
				fmt.Printf("%s withdrew %d. Remaining balance: %d\n", v.name, amount, v.balance)
			} else {
				fmt.Printf("%s has insufficient balance to withdraw %d.\n", v.name, amount)
			}
		case *KazPostAccount:
			if v.balance >= amount {
				v.balance -= amount
				fmt.Printf("%s withdrew %d. Remaining balance: %d\n", v.name, amount, v.balance)
			} else {
				fmt.Printf("%s has insufficient balance to withdraw %d.\n", v.name, amount)
			}
		}
	}
}
func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		switch v := account.(type) {
		case *FedexAccount:
			packageDetails := fmt.Sprintf("%s send package to %s", v.name, recipient)
			v.packages = append(v.packages, packageDetails)
			fmt.Println(packageDetails)
		case *KazPostAccount:
			packageDetails := fmt.Sprintf("%s send package to %s", v.name, recipient)
			v.packages = append(v.packages, packageDetails)
			fmt.Println(packageDetails)
		}
	}
}
