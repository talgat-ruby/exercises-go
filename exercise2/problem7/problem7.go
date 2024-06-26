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

type post interface {
	sendPackage(name string)
}

type bank interface {
	withdraw(amount int)
}

func (k *KazPostAccount) sendPackage(Name string) {
	k.packages = append(k.packages, fmt.Sprintf("%s send package to %s", k.name, Name))
}

func (f *FedexAccount) sendPackage(Name string) {
	f.packages = append(f.packages, fmt.Sprintf("%s send package to %s", f.name, Name))
}

func (b *BankAccount) withdraw(amount int) {
	if b.balance < amount {
		fmt.Printf("Not enough balance in %v\n", b.name)
	} else {
		b.balance -= amount
	}
}

func (k *KazPostAccount) withdraw(amount int) {
	if k.balance < amount {
		fmt.Printf("Not enough balance in %v\n", k.name)
	} else {
		k.balance -= amount
	}
}

func withdrawMoney(amount int, accounts ...bank) {
	for _, account := range accounts {
		account.withdraw(amount)
	}
}

func sendPackagesTo(name string, senders ...post) {
	for _, sebder := range senders {
		sebder.sendPackage(name)
	}
}
