package problem3

func sum(a, b int) int {
	var c int
	ch := make(chan int)

	go func(a, b int, ch chan int) {
		c = a + b
		ch <- c
	}(a, b, ch)

	<-ch
	return c
}
