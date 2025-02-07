// File: problem2.go
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
	numCores := runtime.NumCPU()
	chunkSize := (len(numbers) + numCores - 1) / numCores
	var sum int64
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	wg.Add(numCores)
	for i := 0; i < numCores; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}

		go func(start, end int) {
			defer wg.Done()
			var localSum int64
			for _, n := range numbers[start:end] {
				localSum += int64(n)
			}
			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()
	return sum
}
