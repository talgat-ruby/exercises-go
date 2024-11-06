package problem4

import "sync"

func iter(ch chan<- int, nums []int, wg *sync.WaitGroup) {
	for _, n := range nums {
		ch <- n
	}
	wg.Done()
}

func sum(nums []int) int {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go iter(ch, nums, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	var sum int
	for n := range ch {
		sum += n
	}
	return sum
}
