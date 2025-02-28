package problem3

func sum(a, b int) int {
	result := make(chan int)

	go func(a, b int) {
		result <- a + b
	}(a, b)

	return <-result
}
