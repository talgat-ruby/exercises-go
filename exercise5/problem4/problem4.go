package problem4

func iter(ch chan<- int, nums []int) {
	for _, n := range nums {
		ch <- n
	}
	close(ch)
}

func sum(nums []int) int {
	ch := make(chan int, len(nums))
	done := make(chan bool)

	go iter(ch, nums)

	var sum int
	go func() {
		for {
			n, more := <-ch
			if more {
				sum += n
			} else {
				done <- true
				break
			}
		}
	}()
	<-done
	return sum
}
