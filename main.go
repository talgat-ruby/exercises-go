package main

import (
	"fmt"
	"sync"
)

type bankAccount struct {
	blnc int
	m    *sync.Mutex
}

func newAccount(blnc int) *bankAccount {
	m := sync.Mutex{}
	return &bankAccount{blnc, &m}
}

func main() {
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

	fmt.Println(acc.blnc)
	if acc.blnc != exp {
		fmt.Println("erroorrrrrr")
	} else {
		fmt.Println("vse ok")
	}
}

func (b *bankAccount) withdraw(balance int) {
	b.m.Lock()
	if b.blnc > 10 {
		b.blnc = b.blnc - balance
	}
	b.m.Unlock()
}

func main() {
	//
}
