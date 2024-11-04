package problem2

import (
	"runtime"
	"sync"
)

func add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

func addConcurrently(numbers []int) int64 {
	var sum int64
	numCPU := runtime.NumCPU()
	chunkSize := len(numbers) / numCPU

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numCPU; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == numCPU-1 {
			end = len(numbers)
		}

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
