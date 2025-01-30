package problem3

func sum(a, b int) int {
	var c int
	ch := make(chan int)
	go func(a, b int, ch chan int) {
		sum := a + b
		ch <- sum
	}(a, b, ch)
	c = <-ch
	return c
}
