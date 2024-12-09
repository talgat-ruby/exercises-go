package problem2

import "runtime"

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
	ch := make(chan int64)
	cpuNum := runtime.NumCPU()

	go addSlcToCh(slc)
	for i := 0; i < cpuNum; i++ {

	}
	return sum
}
