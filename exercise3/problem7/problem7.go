package problem7

import "fmt"

type Withdrawable interface {
	withdraw(money int)
}

type PackageSendable interface {
	sendPackage(pkg string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) withdraw(money int) {
	b.balance -= money
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) sendPackage(receiver string) {
	msg := fmt.Sprintf("%s send package to %s", f.name, receiver)
	f.packages = append(f.packages, msg)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) withdraw(money int) {
	k.balance -= money
}

func (k *KazPostAccount) sendPackage(receiver string) {
	msg := fmt.Sprintf("%s send package to %s", k.name, receiver)
	k.packages = append(k.packages, msg)
}

func withdrawMoney(money int, accounts ...Withdrawable) {
	for _, account := range accounts {
		account.withdraw(money)
	}
}

func sendPackagesTo(receiver string, accounts ...PackageSendable) {
	for _, account := range accounts {
		account.sendPackage(receiver)
	}
}
