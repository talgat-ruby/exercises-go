package problem2

import (
	"time"
)

var readDelay = 10 * time.Millisecond

type bankAccount struct {
	blnc int
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc}
}

func (b *bankAccount) balance() int {
	time.Sleep(readDelay)
	return 0
}
