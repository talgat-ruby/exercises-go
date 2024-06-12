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

func (b *BankAccount) WithdrawMoney(amount int) {
	b.balance -= amount
}

func (k *KazPostAccount) WithdrawMoney(amount int) {
	k.balance -= amount
}

func (f *FedexAccount) SendPackageTo(receiver string) {
	message := fmt.Sprintf("%s send package to %s", f.name, receiver)
	f.packages = append(f.packages, message)
}

func (k *KazPostAccount) SendPackageTo(receiver string) {
	message := fmt.Sprintf("%s send package to %s", k.name, receiver)
	k.packages = append(k.packages, message)
}

func withdrawMoney(amount int, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *BankAccount:
			acc.WithdrawMoney(amount)
		case *KazPostAccount:
			acc.WithdrawMoney(amount)
		}
	}
}

func sendPackagesTo(receiver string, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *FedexAccount:
			acc.SendPackageTo(receiver)
		case *KazPostAccount:
			acc.SendPackageTo(receiver)
		}
	}
}
