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

type Bank interface {
	withdraw(int)
}

type Post interface {
	send(string)
}

func (b *BankAccount) withdraw(amount int) {
	if b.balance >= amount {
		b.balance -= amount
	}
}

func (k *KazPostAccount) withdraw(amount int) {
	if k.balance >= amount {
		k.balance -= amount
	}
}

func (f *FedexAccount) send(pckg string) {
	msg := fmt.Sprintf("%s send package to %s", f.name, pckg)
	f.packages = append(f.packages, msg)
}

func (k *KazPostAccount) send(pckg string) {
	msg := fmt.Sprintf("%s send package to %s", k.name, pckg)
	k.packages = append(k.packages, msg)
}

func withdrawMoney(amount int, people ...Bank) {
	for _, person := range people {
		person.withdraw(amount)
	}
}

func sendPackagesTo(pckg string, people ...Post) {
	for _, person := range people {
		person.send(pckg)
	}
}
