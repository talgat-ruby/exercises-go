package problem3

func sum(a, b int) int {
	channel := make(chan int)
	go func(a, b int) {
		channel <- a + b
		close(channel)
	}(a, b)

	return <-channel
}
