package problem2

import (
	"runtime"
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
	numCPUs := runtime.NumCPU()
	chunkSize := len(numbers) / numCPUs

	var wg sync.WaitGroup
	wg.Add(numCPUs)

	for i := 0; i < numCPUs; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numCPUs-1 {
			end = len(numbers)
		}

		go func(nums []int) {
			defer wg.Done()

			var partialSum int64
			for _, n := range nums {
				partialSum += int64(n)
			}

			atomic.AddInt64(&sum, partialSum)
		}(numbers[start:end])
	}

	wg.Wait()
	return sum
}
