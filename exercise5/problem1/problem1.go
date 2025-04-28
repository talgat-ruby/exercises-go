package problem1

func incrementConcurrently(num int) int {
	ch := make(chan int)
	go func() {
		num++
		ch <- num
	}()
	<-ch
	return num
}
