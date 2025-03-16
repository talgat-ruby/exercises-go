package main

func addUp(x int) int {
	var c int = 0
	var i int
	for i = 1; i <= x; i++ {
		c += i
	}
	return c
}
