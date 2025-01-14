package problem1

func incrementConcurrently(num int) int {
	ch := make(chan int)
	defer close(ch)
	go func() {
		num++
		ch <- 1
	}()
	<-ch
	return num
}
