package problem1

import "sync"

func incrementConcurrently(num int) int {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		num++
	}()

	wg.Wait()
	return num
}
