package problem7

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

func withdrawMoney(amount int, people ...interface{}) {
	for _, person := range people {
		switch v := person.(type) {
		case BankAccount:
			if v.balance >= amount {
				v.balance -= amount
			}
		case KazPostAccount:
			if v.balance >= amount {
				v.balance -= amount
			}
		}
	}
}

func sendPackagesTo(pckg string, people ...interface{}) {
	for _, person := range people {
		switch v := person.(type) {
		case FedexAccount:
			v.packages = append(v.packages, pckg)
		case KazPostAccount:
			v.packages = append(v.packages, pckg)
		}
	}
}
