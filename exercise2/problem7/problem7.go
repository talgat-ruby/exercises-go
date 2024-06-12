package problem7

import "errors"

type Account interface {
	Withdraw(amount int) error
}

type PackageSender interface {
	SendPackageTo(name string) error
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount > b.balance {
		return errors.New("not enough money")
	}
	b.balance -= amount
	return nil
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) SendPackageTo(name string) error {
	f.packages = append(f.packages, f.name+" send package to "+name)
	return nil
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) Withdraw(amount int) error {
	if amount > k.balance {
		return errors.New("not enough money")
	}
	k.balance -= amount
	return nil
}

func (k *KazPostAccount) SendPackageTo(name string) error {
	k.packages = append(k.packages, k.name+" send package to "+name)
	return nil
}

func withdrawMoney(amount int, accounts ...Account) {
	for _, a := range accounts {
		a.Withdraw(amount)
	}
}

func sendPackagesTo(name string, accounts ...PackageSender) {
	for _, a := range accounts {
		a.SendPackageTo(name)
	}
}
