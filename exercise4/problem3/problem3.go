package problem3

func sum(a, b int) int {
	var c int
	result := make(chan int)

	go func(a, b int) {
		c = a + b
		result <- c
	}(a, b)

	return <-result
}
