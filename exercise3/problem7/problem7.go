package problem7

import "fmt"

type Withdrawer interface {
	Withdraw(amount int) error
	GetBalance() int
}

type PackageSender interface {
	SendPackage(receiver string)
	GetPackages() []string
}

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

func (ba *BankAccount) Withdraw(amount int) error {
	if ba.balance < amount {
		return fmt.Errorf("%s has insufficient balance", ba.name)
	}
	ba.balance -= amount
	return nil
}

func (ka *KazPostAccount) Withdraw(amount int) error {
	if ka.balance < amount {
		return fmt.Errorf("%s has insufficient balance", ka.name)
	}
	ka.balance -= amount
	return nil
}

func (fa *FedexAccount) SendPackage(receiver string) {
	fa.packages = append(fa.packages, fa.name+" send package to "+receiver)
	fmt.Println(fa.name + " send package to " + receiver)
}

func (ka *KazPostAccount) SendPackage(receiver string) {
	ka.packages = append(ka.packages, ka.name+" send package to "+receiver)
	fmt.Println(ka.name + " send package to " + receiver)
}

func (ba *BankAccount) GetBalance() int {
	return ba.balance
}

func (ka *KazPostAccount) GetBalance() int {
	return ka.balance
}

func (ka *KazPostAccount) GetPackages() []string {
	return ka.packages
}

func (fa *FedexAccount) GetPackages() []string {
	return fa.packages
}

func withdrawMoney(amount int, withdrawers ...Withdrawer) {
	for _, account := range withdrawers {
		err := account.Withdraw(amount)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T: Withdrawn %d, remaining balance: %d\n", account, amount, account.GetBalance())
		}
	}
}

func sendPackagesTo(receiver string, senders ...PackageSender) {
	for _, sender := range senders {
		sender.SendPackage(receiver)
	}
}
