package problem1

func incrementConcurrently(num int) int {
	c := make(chan int)
	go func() {
		c <- num + 1
		close(c)
	}()
	return <-c
}
