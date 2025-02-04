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
	runtime.GOMAXPROCS(numCores)

	chunkSize := len(numbers) / numCores
	results := make([]int64, numCores)

	var wg sync.WaitGroup
	var sum int64

	for i := 0; i < numCores; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == numCores-1 {
			end = len(numbers)
		}

		go func(i, start, end int) {
			defer wg.Done()
			results[i] = add(numbers[start:end])
		}(i, start, end)
	}

	wg.Wait()

	for _, s := range results {
		sum += s
	}
	return sum
}
