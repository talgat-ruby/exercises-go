package problem7

type Account interface {
	getName() string
	getBalance() int
	withdraw(newBalance int)
}

type Sender interface {
	getName() string
	sendPackage(recipient string)
}

func withdrawMoney(value int, accounts ...Account) {
	for _, account := range accounts {
		account.withdraw(value)
	}
}

func sendPackagesTo(recipient string, sender ...Sender) {
	for _, sender := range sender {
		sender.sendPackage(recipient)
	}
}
