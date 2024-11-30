package problem3

func sum(a, b int) int {
	var c int
	ch := make(chan bool)
	go func(a, b int) {
		c = a + b
		ch <- true
		close(ch)
	}(a, b)
	<-ch
	return c
}
