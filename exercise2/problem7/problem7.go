package problem7

import "fmt"

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

type withDrawable interface {
	withDraw(int) int
}

type deliverable interface {
	delivery(string)
}

func (b *BankAccount) withDraw(amount int) int {
	b.balance -= amount

	return b.balance
}

func (post *KazPostAccount) withDraw(amount int) int {
	post.balance -= amount

	return post.balance
}

func (fedex *FedexAccount) delivery(to string) {
	fedex.packages = append(fedex.packages, fmt.Sprintf("%s send package to %s", fedex.name, to))
}

func (post *KazPostAccount) delivery(to string) {
	post.packages = append(post.packages, fmt.Sprintf("%s send package to %s", post.name, to))
}

func withdrawMoney(amount int, drawableAccounts ...withDrawable) {
	for _, drawableAccount := range drawableAccounts {
		drawableAccount.withDraw(amount)
	}
}

func sendPackagesTo(to string, senders ...deliverable) {
	for _, sender := range senders {
		sender.delivery(to)
	}
}
