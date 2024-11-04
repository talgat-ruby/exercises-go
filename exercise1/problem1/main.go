package main

func addUp(a int) int {
	sum := 0
	for i := 0; i <= a; i++ {
		sum += i
	}
	return sum
}
