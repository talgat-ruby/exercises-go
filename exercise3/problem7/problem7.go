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

func (b *BankAccount) getName() string {
	return b.name
}

func (b *FedexAccount) getName() string {
	return b.name
}

func (b *KazPostAccount) getName() string {
	return b.name
}

func (b *BankAccount) withdraw(amount int) {
	b.balance -= amount
}

func (b *KazPostAccount) withdraw(amount int) {
	b.balance -= amount
}

func (b *FedexAccount) send(receiver string, name string) {
	b.packages = append(b.packages, name+" send package to "+receiver)
}

func (b *KazPostAccount) send(receiver string, name string) {
	b.packages = append(b.packages, name+" send package to "+receiver)
}

type HasBalance interface {
	withdraw(amount int)
}

type HasPackages interface {
	send(receiver string, name string)
	getName() string
}

func withdrawMoney(amount int, people ...HasBalance) {
	for _, p := range people {
		p.withdraw(amount)
	}
}

func sendPackagesTo(receiver string, people ...HasPackages) {
	for _, p := range people {
		p.send(receiver, p.getName())
	}
}
