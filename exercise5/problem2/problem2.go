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

// addConcurrently divides the work among available CPU cores and adds the numbers concurrently.
func addConcurrently(numbers []int) int64 {
	var sum int64
	numCores := runtime.NumCPU()
	chunkSize := (len(numbers) + numCores - 1) / numCores
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numCores; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}

		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			var localSum int64
			for _, n := range nums {
				localSum += int64(n)
			}
			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait()
	return sum
}
