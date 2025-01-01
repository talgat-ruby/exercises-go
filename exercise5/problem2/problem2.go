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
	partSize := len(numbers) / numCores
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numCores; i++ {
		wg.Add(1)

		start := i * partSize
		end := start + partSize
		if i == numCores-1 {
			end = len(numbers)
		}
		go func(part []int) {
			defer wg.Done()
			var partSum int64
			for _, n := range part {
				partSum += int64(n)
			}
			mu.Lock()
			sum += partSum
			mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait()
	return sum
}
