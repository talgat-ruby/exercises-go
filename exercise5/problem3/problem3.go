package problem3

func sum(a, b int) int {
	resultChan := make(chan int)
	go func(a, b int) {
		resultChan <- a + b
	}(a, b)
	return <-resultChan
}
