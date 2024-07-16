package problem3

func sum(a, b int) int {
	c := make(chan int)

	go func(a, b int, c chan<- int) {
		c <- a + b
	}(a, b, c)

	result := <-c
	return result
}
