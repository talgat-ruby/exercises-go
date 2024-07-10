package problem3

func sum(a, b int) int {
	var c int

	go func(a, b int) {
		c = a + b
	}(a, b)

	return c
}
