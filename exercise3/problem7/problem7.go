package problem7

import "fmt"

type Treasury interface {
	withdraw(int)
}

type Post interface {
	sendTo(string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) withdraw(amount int) {
	b.balance -= amount
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) sendTo(name string) {
	f.packages = append(f.packages, fmt.Sprintf("%s send package to %s", f.name, name))
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) withdraw(amount int) {
	k.balance -= amount
}

func (k *KazPostAccount) sendTo(name string) {
	k.packages = append(k.packages, fmt.Sprintf("%s send package to %s", k.name, name))
}

func withdrawMoney(amount int, accounts ...Treasury) {
	for _, account := range accounts {
		account.withdraw(amount)
	}
}
func sendPackagesTo(name string, posts ...Post) {
	for _, post := range posts {
		post.sendTo(name)
	}
}
