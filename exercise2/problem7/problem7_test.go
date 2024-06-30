package problem7

type hasBalance interface {
	Withdraw(amount int)
}

type BankAccount struct {
	name    string
	balance int
	hasBalance
}

type FedexAccount struct {
	name     string
	packages []string
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
	hasBalance
}

func (acc *BankAccount) Withdraw(amount int) {
	if amount > 0 && amount <= acc.balance {
		acc.balance -= amount
	}
}

func (acc *KazPostAccount) Withdraw(amount int) {
	if amount > 0 && amount <= acc.balance {
		acc.balance -= amount
	}
}

func (acc *FedexAccount) AddPackage(pkg string) {
	acc.packages = append(acc.packages, pkg)
}

func (acc *KazPostAccount) AddPackage(pkg string) {
	acc.packages = append(acc.packages, pkg)
}

func withdrawMoney(amount int, accounts ...any) {
	for _, account := range accounts {
		switch account.(type) {
		case *BankAccount:
			account.(*BankAccount).Withdraw(amount)
		case *KazPostAccount:
			account.(*KazPostAccount).Withdraw(amount)
		}
	}
}

func sendPackagesTo(pkg string, accounts ...any) {
	postfix := " send package to " + pkg
	for _, account := range accounts {
		if acc, ok := account.(*FedexAccount); ok {
			acc.AddPackage(acc.name + postfix)
		}
		if acc, ok := account.(*KazPostAccount); ok {
			acc.AddPackage(acc.name + postfix)
		}
	}
}
