package problem2

import (
	"runtime"
	"sync"
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

	pCount := runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(pCount)
	items := len(numbers) / pCount

	for i := range pCount {
		part := numbers[i*items : (i+1)*items]
		go func(part []int) {
			defer wg.Done()
			sum += add(part)
		}(part)
	}

	wg.Wait()

	return sum
}
