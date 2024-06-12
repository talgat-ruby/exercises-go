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

type bank interface {
	withdrawn(int)
}

type post interface {
	sent(string)
}

func (b *BankAccount) withdrawn(amount int) {
	b.balance -= amount
}

func (k *KazPostAccount) withdrawn(amount int) {
	k.balance -= amount
}

func (f *FedexAccount) sent(to string) {
	f.packages = append(f.packages, fmt.Sprintf("%s send package to %s", f.name, to))
}

func (k *KazPostAccount) sent(to string) {
	k.packages = append(k.packages, fmt.Sprintf("%s send package to %s", k.name, to))
}

func withdrawMoney(amount int, clients ...bank) {
	for _, client := range clients {
		client.withdrawn(amount)
	}
}

func sendPackagesTo(to string, senders ...post) {
	for _, sender := range senders {
		sender.sent(to)
	}
}
