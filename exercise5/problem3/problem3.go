package problem3

func sum(a, b int) int {
	var c int
	ch := make(chan struct{})
	go func(a, b int) {
		c = a + b
		close(ch)
	}(a, b)
	<-ch
	return c
}



