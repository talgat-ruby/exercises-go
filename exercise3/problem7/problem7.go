package problem7

type Pocket interface {
	moneyLeft(x int) int
}

type Box interface {
	insideBox(name string) []string
}

type BankAccount struct {
	name    string
	balance int
}

func (b *BankAccount) moneyLeft(x int) int {
	b.balance = b.balance - x
	return b.balance
}

type FedexAccount struct {
	name     string
	packages []string
}

func (f *FedexAccount) insideBox(name string) []string {
	s := f.name + " send package to " + name
	f.packages = append(f.packages, s)
	return f.packages
}

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (k *KazPostAccount) insideBox(name string) []string {
	s := k.name + " send package to " + name
	k.packages = append(k.packages, s)
	return k.packages
}

func (k *KazPostAccount) moneyLeft(x int) int {
	k.balance = k.balance - x
	return k.balance
}

func withdrawMoney(amount int, pockets ...Pocket) {
	for _, pocket := range pockets {
		pocket.moneyLeft(amount)
	}
}

func sendPackagesTo(person string, packages ...Box) {
	for _, p := range packages {
		p.insideBox(person)
	}
}
