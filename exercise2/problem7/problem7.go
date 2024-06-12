package problem7

import "errors"

type Withdrawable interface {
	WithdrawMoney(amount int) error
	RecieveMoney(amount int) error
}

type Sendable interface {
	RecievePackage(packageName string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b BankAccount) WithdrawMoney(amount int) error {
	if b.balance < amount {
		return errors.New("Not enough money")
	}
	b.balance -= amount
	return nil
}

type FedexAccount struct {
	name     string
	Balance  int
	packages []string
}

func (f FedexAccount) RecievePackage(packageName string) {
	f.packages = append(f.packages, packageName)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k KazPostAccount) RecievePackage(packageName string) {
	k.packages = append(k.packages, packageName)
}

func (k KazPostAccount) WithdrawMoney(amount int) error {
	if k.balance < amount {
		return errors.New("Not enough money")
	}
	k.balance -= amount
	return nil
}

func (b BankAccount) RecieveMoney(amount int) error {
	b.balance += amount
	return nil
}

func (k KazPostAccount) RecieveMoney(amount int) error {
	k.balance += amount
	return nil
}

func withdrawMoney(amount int, w, r Withdrawable) {
	w.WithdrawMoney(amount)
	r.RecieveMoney(amount)
}

func sendPackagesTo(packageName string, recievers ...Sendable) {
	for _, r := range recievers {
		r.RecievePackage(packageName)
	}
}
