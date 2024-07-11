package problem4

func iter(ch chan<- int, nums []int, done chan<- bool) {
	for _, n := range nums {
		ch <- n
	}
	done <- true
}

func sum(nums []int) int {
	ch := make(chan int, len(nums))
	done := make(chan bool)

	go iter(ch, nums, done)
	<-done
	close(done)
	close(ch)

	var sum int
	for v := range ch {
		sum += v
	}
	return sum
}
