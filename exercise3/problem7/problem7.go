package problem7

import "fmt"

type Drawable interface {
	transaction(sum int)
}

type Send interface {
	sendPackage(name string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) transaction(sum int) {
	b.balance = b.balance - sum
}

type FedexAccount struct {
	name     string
	packages []string
}

func (s *FedexAccount) sendPackage(name string) {
	s.packages = append(s.packages, fmt.Sprintf("%s send package to %s", s.name, name))
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (b *KazPostAccount) transaction(sum int) {
	b.balance = b.balance - sum
}

func (s *KazPostAccount) sendPackage(name string) {
	s.packages = append(s.packages, fmt.Sprintf("%s send package to %s", s.name, name))
}

func withdrawMoney(sum int, accounts ...Drawable) {
	for _, v := range accounts {
		v.transaction(sum)
	}
}

func sendPackagesTo(name string, accounts ...Send) {
	for _, v := range accounts {
		v.sendPackage(name)
	}
}
