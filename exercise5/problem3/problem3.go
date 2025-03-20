package problem3

func sum(a, b int) int {
	channel := make(chan int)

	go func(a, b int) {
		sum := a + b
		channel <- sum
	}(a, b)
	c := <-channel
	return c
}
