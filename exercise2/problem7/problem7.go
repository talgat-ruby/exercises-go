package problem7

type Withdrawable interface {
	GetBalance() int
	SetBalance(int)
}

type Sendable interface {
	GetName() string
	GetPackages() []string
	SetPackages([]string)
}

type BankAccount struct {
	name     string
	balance  int
	packages []string
}

type FedexAccount struct {
	name     string
	balance  int
	packages []string
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (b *BankAccount) GetName() string {
	return b.name
}

func (b *FedexAccount) GetName() string {
	return b.name
}

func (b *KazPostAccount) GetName() string {
	return b.name
}

func (b *BankAccount) GetBalance() int {
	return b.balance
}

func (b *FedexAccount) GetBalance() int {
	return b.balance
}

func (b *KazPostAccount) GetBalance() int {
	return b.balance
}

func (b *BankAccount) SetBalance(i int) {
	b.balance = i
}

func (b *FedexAccount) SetBalance(i int) {
	b.balance = i
}

func (b *KazPostAccount) SetBalance(i int) {
	b.balance = i
}

func (b *BankAccount) GetPackages() []string {
	return b.packages
}

func (b *FedexAccount) GetPackages() []string {
	return b.packages
}

func (b *KazPostAccount) GetPackages() []string {
	return b.packages
}

func (b *BankAccount) SetPackages(p []string) {
	b.packages = p
}

func (b *FedexAccount) SetPackages(p []string) {
	b.packages = p
}

func (b *KazPostAccount) SetPackages(p []string) {
	b.packages = p
}

func withdrawMoney(m int, a ...Withdrawable) {
	for _, account := range a {
		account.SetBalance(account.GetBalance() - m)
	}
}

func sendPackagesTo(name string, s ...Sendable) {
	for _, sender := range s {
		sender.SetPackages(append(sender.GetPackages(), sender.GetName()+" send package to "+name))
	}
}
