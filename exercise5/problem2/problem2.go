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
	cpuCount := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuCount)
	chunkSize := len(numbers) / cpuCount
	resultCh := make(chan int, cpuCount)
	wg := new(sync.WaitGroup)

	var sum int64

	for i := 0; i < cpuCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == cpuCount-1 {
			end = len(numbers)
		}
		wg.Add(1)
		go processAddChunks(numbers[start:end], resultCh, wg)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		sum += int64(result)
	}

	return sum
}

func processAddChunks(chunk []int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for _, num := range chunk {
		result += num
	}
	resultCh <- result
}
