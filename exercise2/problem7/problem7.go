package problem7

type BankAccount struct {
	name    string
	balance float64
}

type FedexAccount struct {
	name     string
	packages []string
}

type KazPostAccount struct {
	name     string
	packages []string
	balance  float64
}

func withdrawMoney(amount int, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *BankAccount:
			acc.balance -= float64(amount)
		case *KazPostAccount:
			acc.balance -= float64(amount)
		}
	}
}

func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *FedexAccount:
			if acc.packages == nil {
				acc.packages = make([]string, 0)
			}
			acc.packages = append(acc.packages, acc.name+" send package to "+recipient)
		case *KazPostAccount:
			if acc.packages == nil {
				acc.packages = make([]string, 0)
			}
			acc.packages = append(acc.packages, acc.name+" send package to "+recipient)
		}
	}
}
