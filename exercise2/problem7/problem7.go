package problem7

type BankAccount struct {
	name    string
	balance int
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

type CanWithdraw interface {
	withdrawMoney(amount int)
}

type CanSendPackages interface {
	sendPackage(to string)
}

func (a *BankAccount) withdrawMoney(amount int) {
	a.balance -= amount
}

func (a *KazPostAccount) withdrawMoney(amount int) {
	a.balance -= amount
}

func (a *FedexAccount) sendPackage(to string) {
	a.packages = make([]string, 1)
	a.packages[0] = a.name + " send package to " + to
}

func (a *KazPostAccount) sendPackage(to string) {
	a.packages = make([]string, 1)
	a.packages[0] = a.name + " send package to " + to
}

func withdrawMoney(amount int, accounts ...CanWithdraw) {
	for _, a := range accounts {
		a.withdrawMoney(amount)
	}
}

func sendPackagesTo(to string, accounts ...CanSendPackages) {
	for _, a := range accounts {
		a.sendPackage(to)
	}
}
