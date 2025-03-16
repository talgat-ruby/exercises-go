package problem7

type withdrawable interface {
	withdraw(sum int)
}

type sendeble interface {
	sendPackages(name string)
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) withdraw(sum int) {
	b.balance -= sum
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) sendPackages(name string) {
	f.packages = append(f.packages, f.name+" send package to "+name)
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) withdraw(sum int) {
	k.balance -= sum
}

func (k *KazPostAccount) sendPackages(name string) {
	k.packages = append(k.packages, k.name+" send package to "+name)
}

func withdrawMoney(sum int, w ...withdrawable) {
	for _, v := range w {
		v.withdraw(sum)
	}
}

func sendPackagesTo(name string, s ...sendeble) {
	for _, v := range s {
		v.sendPackages(name)
	}
}
