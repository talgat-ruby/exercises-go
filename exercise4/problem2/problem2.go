package problem2

import (
	"sync"
	"sync/atomic"
)

// add - sequential code to add numbers, don't update it, just to illustrate concept
func add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

func addConcurrently(numbers []int) int64 {
	var sum int64
	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			atomic.AddInt64(&sum, int64(n))
		}(n)
	}
	wg.Wait()
	return sum
}
