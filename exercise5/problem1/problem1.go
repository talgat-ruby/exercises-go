package problem1

func incrementConcurrently(num int) int {
	ch := make(chan int)
	go func(ch chan<- int) {
		num++
		ch <- num
	}(ch)

	<-ch

	return num
}
