package problem7

import "fmt"

type KazPostAccount struct {
	name     string
	balance  int
	packages []string
}

func (kazPostAcc *KazPostAccount) getName() string {
	return kazPostAcc.name
}

func (kazPostAcc *KazPostAccount) getBalance() int {
	return kazPostAcc.balance
}

func (kazPostAcc *KazPostAccount) withdraw(amount int) {
	kazPostAcc.balance = kazPostAcc.getBalance() - amount
}

func (kazPostAcc *KazPostAccount) sendPackage(recipient string) {
	message := fmt.Sprintf("%s send package to %s", kazPostAcc.getName(), recipient)
	kazPostAcc.packages = append(kazPostAcc.packages, message)
}
