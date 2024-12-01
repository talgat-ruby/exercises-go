package problem3

func sum(a, b int) int {
	ch := make(chan int)

	go func(a, b int) {
		ch <- a + b
	}(a, b)

	return <-ch
}
