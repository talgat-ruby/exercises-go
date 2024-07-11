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
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	var sum int64
	out := make(chan int64, numCpu)

	var wg sync.WaitGroup
	wg.Add(numCpu)

	length := len(numbers)/numCpu + 1
	for i := 0; i < numCpu; i++ {
		minI := i * length
		maxI := min((i+1)*length, len(numbers))
		slice := numbers[minI:maxI]

		go func() {
			out <- add(slice)
			wg.Done()
		}()
	}
	wg.Wait()
	close(out)

	for v := range out {
		sum += v
	}
	return sum
}
