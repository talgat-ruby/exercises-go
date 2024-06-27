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

type bank interface {
	withdrawn(int)
}

type post interface {
	sent(string)
}

func (b *BankAccount) withdrawn(amount int) {
	b.balance -= amount
}

func (k *KazPostAccount) withdrawn(amount int) {
	k.balance -= amount
}

func (f *FedexAccount) sent(receiver string) {
	f.packages = append(f.packages, f.name+" send package to "+receiver)
}

func (k *KazPostAccount) sent(receiver string) {
	k.packages = append(k.packages, k.name+" send package to "+receiver)
}

func withdrawMoney(amount int, clients ...bank) {
	for _, client := range clients {
		client.withdrawn(amount)
	}
}

func sendPackagesTo(to string, senders ...post) {
	for _, sender := range senders {
		sender.sent(to)
	}
}
