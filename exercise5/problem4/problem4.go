package problem4

func iter(ch chan<- int, nums []int) {
	for _, n := range nums {
		ch <- n
	}
}

func sum(nums []int) int {
	ch := make(chan int)

	go iter(ch, nums)

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += <-ch
	}

	return sum
}
