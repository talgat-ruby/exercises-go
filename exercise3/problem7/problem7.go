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
	name string
	balance int
	packages []string
}

func withdrawMoney(num int, client1 *BankAccount, client2 *KazPostAccount) {
	client1.balance -= num
	client2.balance -= num
}

func sendPackagesTo(name string, client1 *FedexAccount, client2 *KazPostAccount) {
	client1.packages = append(client1.packages, client1.name + " send package to " + name)
	client2.packages = append(client2.packages, client2.name + " send package to " + name)
}
