package problem7

type BankAccount struct {
	name    string
	balance int
}

func (bankAcc *BankAccount) getName() string {
	return bankAcc.name
}

func (bankAcc *BankAccount) getBalance() int {
	return bankAcc.balance
}

func (bankAcc *BankAccount) withdraw(amount int) {
	bankAcc.balance = bankAcc.getBalance() - amount
}
