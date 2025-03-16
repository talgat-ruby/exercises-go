package problem3

func sum(a, b int) int {
	var c int
	ch := make(chan bool)

	go func(a, b int, ch chan bool) {
		c = a + b
		ch <- true
	}(a, b, ch)

	<-ch

	return c
}
