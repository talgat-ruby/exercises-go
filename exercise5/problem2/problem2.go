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
	numCPU := runtime.NumCPU()
	countCPU := len(numbers) / numCPU
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sum int64

	for i := 0; i < numCPU; i++ {
		start := i * countCPU
		end := start + countCPU
		if i == numCPU-1 {
			end = len(numbers)
		}

		wg.Add(1)
		go addSum(numbers[start:end], &wg, &mu, &sum)
	}

	wg.Wait()
	return sum
}

func addSum(nums []int, wg *sync.WaitGroup, mu *sync.Mutex, sum *int64) {
	defer wg.Done()
	var s int64
	for _, n := range nums {
		s += int64(n)
	}
	mu.Lock()
	*sum += s
	mu.Unlock()
}
