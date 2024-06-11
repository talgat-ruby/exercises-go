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

type Withdrawable interface {
	withdrawMoney(amount int) error
}

type Sendable interface {
	sendPackagesTo(address string)
}

func (b *BankAccount) withdrawMoney(amount int) error {
	if b.balance > amount {
		b.balance -= amount
	}

	return nil
}

func (k *KazPostAccount) withdrawMoney(amount int) error {
	if k.balance > amount {
		k.balance -= amount
	}

	return nil
}

func (k *KazPostAccount) sendPackagesTo(address string) {
	message := fmt.Sprintf("%s send package to %s", k.name, address)
	k.packages = append(k.packages, message)
}

func (f *FedexAccount) sendPackagesTo(address string) {
	message := fmt.Sprintf("%s send package to %s", f.name, address)
	f.packages = append(f.packages, message)
}

func withdrawMoney(amount int, accounts ...Withdrawable) error {
	for _, account := range accounts {
		err := account.withdrawMoney(amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func sendPackagesTo(address string, senders ...Sendable) {
	for _, sender := range senders {
		sender.sendPackagesTo(address)
	}
}
