package problem1

import (
	"sync"
	"testing"
)

func TestBankAccount(t *testing.T) {
	t.Run(
		"cannot withdraw more than balance", func(t *testing.T) {
			exp := 3
			var wg sync.WaitGroup
			acc := newAccount(43)

			n := 100

			wg.Add(n)
			for i := range n {
				go func(i int) {
					defer wg.Done()
					acc.withdraw(10)
				}(i)
			}

			wg.Wait()

			if acc.blnc != exp {
				t.Errorf("balance was incorrect, expected: %d, got: %d.", exp, acc.blnc)
			}
		},
	)

	t.Run(
		"concurrently modify balance", func(t *testing.T) {
			var wg sync.WaitGroup
			acc := newAccount(0)

			depositNum := 200
			wg.Add(depositNum)
			for i := range depositNum {
				go func(i int) {
					defer wg.Done()
					acc.deposit(1)
				}(i)
			}

			withdrawNum := 100
			wg.Add(withdrawNum)
			for i := range withdrawNum {
				go func(i int) {
					defer wg.Done()
					acc.withdraw(1)
				}(i)
			}

			wg.Wait()

			if acc.blnc != depositNum-withdrawNum {
				t.Errorf("balance was incorrect, expected: %d, got: %d.", depositNum-withdrawNum, acc.blnc)
			}
		},
	)
}
