package problem7

import (
	"testing"
)

func TestAccounts(t *testing.T) {
	normanOsborne := &BankAccount{
		name:    "Norman Osborne",
		balance: 1_000_000,
	}
	peterParker := &FedexAccount{
		name: "Peter Parker",
	}
	auntMay := &KazPostAccount{
		name:    "Aunt May",
		balance: 200,
	}

	withdrawMoney(10, normanOsborne, auntMay)
	sendPackagesTo("Mary Jane", peterParker, auntMay)

	if normanOsborne.balance != 1_000_000-10 {
		t.Errorf("withdrawMoney() was incorrect, got: %d, expected: %d", normanOsborne.balance, 1_000_000-10)
	}
	if auntMay.balance != 200-10 {
		t.Errorf("withdrawMoney() was incorrect, got: %d, expected: %d", auntMay.balance, 200-10)
	}

	if peterParker.packages[0] != "Peter Parker send package to Mary Jane" {
		t.Errorf("sendPackagesTo() was incorrect, got: %s, expected: %s", peterParker.packages[0], "Peter Parker send package to Mary Jane")
	}
	if auntMay.packages[0] != "Aunt May send package to Mary Jane" {
		t.Errorf("sendPackagesTo() was incorrect, got: %s, expected: %s", peterParker.packages[0], "Aunt May send package to Mary Jane")
	}
}
