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

func (k *KazPostAccount) sendPackage(Name string) {
	k.packages = append(k.packages, fmt.Sprintf("%s send package to %s", k.name, Name))
}

func (f *FedexAccount) sendPackage(Name string) {
	f.packages = append(f.packages, fmt.Sprintf("%s send package to %s", f.name, Name))
}

type Package interface {
	sendPackage(name string)
}

type Account interface {
	withdraw(amount int)
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

func withdrawMoney(amount int, accounts ...Account) {
	for _, account := range accounts {
		account.withdraw(amount)
	}
}

func sendPackagesTo(name string, accounts ...Package) {
	for _, account := range accounts {
		account.sendPackage(name)
	}
}
