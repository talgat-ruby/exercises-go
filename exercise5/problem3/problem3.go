package problem3

func sum(a, b int) int {
	// var c int
	c := make(chan int)

	go func(a, b int) {
		c <- a + b
	}(a, b)

	return <-c
}
