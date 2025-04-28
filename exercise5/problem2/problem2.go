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

// Did this example with the help of ChatGPT
func addConcurrently(numbers []int) int64 {
	var sum int64
	numCores := runtime.NumCPU()
	partSize := len(numbers) / numCores
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numCores; i++ {
		start := i * partSize
		end := start + partSize
		if i == numCores-1 {
			end = len(numbers)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var partialSum int64
			for _, n := range numbers[start:end] {
				partialSum += int64(n)
			}

			mu.Lock()
			sum += partialSum
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()
	return sum
}
