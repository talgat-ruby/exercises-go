package problem3

func sum(a, b int) int {
	ch := make(chan struct{})
	var c int

	go func(a, b int) {
		c = a + b
		ch <- struct{}{}
	}(a, b)

	<-ch

	return c
}
