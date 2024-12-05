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
	quantityOfCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(quantityOfCpu)

	sizeOfBlock := (len(numbers) + quantityOfCpu - 1) / quantityOfCpu
	var wg sync.WaitGroup
	result := make([]int64, quantityOfCpu)

	for i := 0; i < quantityOfCpu; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * sizeOfBlock
			end := start + sizeOfBlock
			if end > len(numbers) {
				end = len(numbers)
			}
			for _, n := range numbers[start:end] {
				result[i] += int64(n)
			}
		}(i)
	}

	wg.Wait()

	var sum int64

	for _, result := range result {
		sum += result
	}

	return sum
}
