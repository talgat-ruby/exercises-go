package problem7

type Withdrawal interface {
	Withdraw(amount int) bool
}

type PackageSender interface {
	SendPackage(to string)
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

func (b *BankAccount) Withdraw(amount int) bool {
	if amount > b.balance {
		return false
	}
	b.balance -= amount
	return true
}

func (k *KazPostAccount) Withdraw(amount int) bool {
	if amount > k.balance {
		return false
	}
	k.balance -= amount
	return true
}

func (account *FedexAccount) SendPackage(to string) {
	message := account.name + " send package to " + to
	account.packages = append(account.packages, message)
}

func (account *KazPostAccount) SendPackage(to string) {
	message := account.name + " send package to " + to
	account.packages = append(account.packages, message)
}

func withdrawMoney(amount int, accounts ...Withdrawal) {
	for _, account := range accounts {
		account.Withdraw(amount)
	}
}

func sendPackagesTo(to string, accounts ...PackageSender) {
	for _, account := range accounts {
		account.SendPackage(to)
	}
}
