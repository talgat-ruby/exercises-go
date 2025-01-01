package problem4

func iter(ch chan<- int, nums []int) {
	for _, n := range nums {
		ch <- n
	}
	close(ch)
}

func sum(nums []int) int {
	ch := make(chan int)

	go iter(ch, nums)

	var total int

	for n := range ch {
		total += n
	}
	return total
}
