package problem7

type Account interface {
	nameAccount() string
}

type WithdrawInterface interface {
	Account
	Withdraw(amount int) bool
}

type SendPackageInterface interface {
	Account
	SendPackage(name string) string
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) nameAccount() string {
	return b.name
}

func (b *BankAccount) Withdraw(money int) bool {
	if b.balance >= money {
		b.balance -= money
		return true
	}
	return false
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) nameAccount() string {
	return f.name
}

func (f *FedexAccount) SendPackage(name string) string {
	message := f.name + " send package to " + name
	f.packages = append(f.packages, message)
	return message
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) nameAccount() string {
	return k.name
}

func (k *KazPostAccount) Withdraw(money int) bool {
	if k.balance >= money {
		k.balance -= money
		return true
	}
	return false
}

func (k *KazPostAccount) SendPackage(name string) string {
	message := k.name + " send package to " + name
	k.packages = append(k.packages, message)
	return message
}

func withdrawMoney(money int, accounts ...WithdrawInterface) {
	for _, account := range accounts {
		account.Withdraw(money)
	}
}

func sendPackagesTo(name string, accounts ...SendPackageInterface) {
	for _, account := range accounts {
		account.SendPackage(name)
	}
}
