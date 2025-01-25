package problem7

type Sender interface {
	send(name string)
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
	packages []string
	balance  int
}
type withdrawable interface {
	withdrow(money int) bool
}

func withdrawMoney(money int, account ...withdrawable) {
	for _, v := range account {
		v.withdrow(money)
	}
}
func sendPackagesTo(to string, fedex *FedexAccount, kaz *KazPostAccount) {
	message := fedex.name + " send package to " + to
	fedex.packages = append(fedex.packages, message)

	kazpostmesage := kaz.name + " send package to " + to
	kaz.packages = append(kaz.packages, kazpostmesage)
}
func (k *KazPostAccount) withdrow(money int) bool {
	if k.balance >= money {
		k.balance -= money
		return true
	}
	return false
}

func (b *BankAccount) withdrow(money int) bool {
	if b.balance >= money {
		b.balance -= money
		return true
	}
	return false
}
