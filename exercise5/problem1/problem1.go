package problem1

import "sync"

func incrementConcurrently(num int) int {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		num++
	}()

	wg.Wait()
	return num
}
