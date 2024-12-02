package problem1

type bankAccount struct {
	blnc int
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc}
}
