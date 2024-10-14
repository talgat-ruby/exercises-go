package problem7

type Withdrawable interface {
	Withdraw(amount int)
}

type PackageSender interface {
	SendPackage(to string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) Withdraw(amount int) {
	b.balance -= amount
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) SendPackage(to string) {
	f.packages = append(f.packages, f.name+" send package to "+to)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (b *KazPostAccount) Withdraw(amount int) {
	b.balance -= amount
}

func (b *KazPostAccount) SendPackage(to string) {
	b.packages = append(b.packages, b.name+" send package to "+to)
}

func withdrawMoney(amount int, accounts ...Withdrawable) {
	for _, account := range accounts {
		account.Withdraw(amount)
	}
}

func sendPackagesTo(to string, senders ...PackageSender) {
	for _, sender := range senders {
		sender.SendPackage(to)
	}
}
