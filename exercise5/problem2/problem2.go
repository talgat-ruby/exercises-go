package problem2

import (
	"runtime"
	"sync"
)

// addConcurrently calculates the sum of the input slice in parallel.
func addConcurrently(nums []int) int {
	numCPUs := runtime.NumCPU()              // Get the number of CPU cores available
	partSize := (len(nums) + numCPUs - 1) / numCPUs // Calculate the size of each part
	results := make([]int, numCPUs)          // Slice to hold partial results
	var wg sync.WaitGroup

	// Launch a goroutine for each segment
	for i := 0; i < numCPUs; i++ {
		wg.Add(1)
		start := i * partSize
		end := start + partSize
		if end > len(nums) {
			end = len(nums)
		}
		go func(i, start, end int) {
			defer wg.Done()
			sum := 0
			for _, num := range nums[start:end] {
				sum += num
			}
			results[i] = sum
		}(i, start, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Aggregate results from all goroutines
	totalSum := 0
	for _, partialSum := range results {
		totalSum += partialSum
	}
	return totalSum
}
