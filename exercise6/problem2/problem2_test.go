package problem2

import (
	"sync"
	"testing"
	"time"
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

			if acc.balance() != exp {
				t.Errorf("balance was incorrect, expected: %d, got: %d.", exp, acc.balance())
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

			if acc.balance() != depositNum-withdrawNum {
				t.Errorf("balance was incorrect, expected: %d, got: %d.", depositNum-withdrawNum, acc.balance())
			}
		},
	)

	t.Run(
		"concurrently read balance", func(t *testing.T) {
			var wg sync.WaitGroup
			acc := newAccount(0)

			readNum := 100

			totalDelay := time.Duration(readNum+10) * readDelay
			timer := time.AfterFunc(
				totalDelay, func() {
					t.Error("too long operation, make it faster please")
				},
			)

			wg.Add(readNum)
			go func() {
				for range readNum {
					go func() {
						defer wg.Done()
						acc.balance()
					}()
				}
			}()

			wg.Add(readNum)
			go func() {
				for range readNum {
					go func() {
						defer wg.Done()
						acc.balance()
					}()
				}
			}()

			wg.Wait()
			timer.Stop()
		},
	)
}
