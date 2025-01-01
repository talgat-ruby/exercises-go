package problem7

type Account interface {
	Withdraw(amount int)
	SendPackage(to string)
}
type BankAccount struct {
	name    string
	balance int
}

func (a *BankAccount) Withdraw(amount int) {
	a.balance -= amount
}
func (a *BankAccount) SendPackage(to string) {

}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) Withdraw(amount int) {

}
func (f *FedexAccount) SendPackage(to string) {
	f.packages = append(f.packages, f.name+" send package to "+to)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) Withdraw(amount int) {
	k.balance -= amount
}
func (k *KazPostAccount) SendPackage(to string) {
	k.packages = append(k.packages, k.name+" send package to "+to)
}

func withdrawMoney(amount int, accounts ...Account) {
	for _, account := range accounts {
		account.Withdraw(amount)
	}
}

func sendPackagesTo(to string, accounts ...Account) {
	for _, account := range accounts {
		account.SendPackage(to)
	}
}
