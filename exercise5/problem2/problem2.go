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
	numCores := runtime.NumCPU()
	chunkSize := (len(numbers) + numCores - 1) / numCores

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < len(numbers); i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}

		go func(chunk []int) {
			defer wg.Done()

			localSum := int64(0)
			for _, n := range chunk {
				localSum += int64(n)
			}

			mu.Lock()
			sum += localSum
			mu.Unlock()
		}(numbers[i:end])
	}
	wg.Wait()
	return sum
}
