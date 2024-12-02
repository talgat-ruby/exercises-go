package problem1

func incrementConcurrently(num int) int {
	ch := make(chan int)
	go func() {
		ch <- num + 1
	}()

	return <-ch
}
