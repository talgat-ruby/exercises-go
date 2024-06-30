package problem7

import "fmt"

type Withdrawable interface {
	Withdraw(int) error
}

type Packaged interface {
	SendPackage(recipient string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) Withdraw(amount int) error {
	if b.balance-amount < 0 {
		return fmt.Errorf("not enough balance")
	}
	b.balance -= amount
	return nil
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) SendPackage(recipient string) {
	message := fmt.Sprintf("%s send package to %s", f.name, recipient)
	f.packages = append(f.packages, message)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) Withdraw(amount int) error {
	if k.balance-amount < 0 {
		return fmt.Errorf("not enough balance")
	}
	k.balance -= amount
	return nil
}

func (k *KazPostAccount) SendPackage(recipient string) {
	message := fmt.Sprintf("%s send package to %s", k.name, recipient)
	k.packages = append(k.packages, message)
}

func withdrawMoney(amount int, accounts ...Withdrawable) error {
	for _, account := range accounts {
		err := account.Withdraw(amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func sendPackagesTo(recipient string, senders ...Packaged) {
	for _, sender := range senders {
		sender.SendPackage(recipient)
	}
}
