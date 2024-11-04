package problem7

import "fmt"

type MoneyWithdrawable interface {
	Withdraw(amount int) error
}

type PackageSendable interface {
	SendPackage(recipient string)
}

type BankAccount struct {
	name    string
	balance int
}

func (ba *BankAccount) Withdraw(amount int) error {
	if ba.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	ba.balance -= amount
	return nil
}

type FedexAccount struct {
	name     string
	packages []string
}

func (fa *FedexAccount) SendPackage(recipient string) {
	message := fmt.Sprintf("%s send package to %s", fa.name, recipient)
	fa.packages = append(fa.packages, message)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (ka *KazPostAccount) Withdraw(amount int) error {
	if ka.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	ka.balance -= amount
	return nil
}

func (ka *KazPostAccount) SendPackage(recipient string) {
	message := fmt.Sprintf("%s send package to %s", ka.name, recipient)
	ka.packages = append(ka.packages, message)
}

func withdrawMoney(amount int, accounts ...MoneyWithdrawable) {
	for _, account := range accounts {
		_ = account.Withdraw(amount)
	}
}

func sendPackagesTo(recipient string, accounts ...PackageSendable) {
	for _, account := range accounts {
		account.SendPackage(recipient)
	}
}
