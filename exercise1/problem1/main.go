package main

func addUp(until int) int {
	var sum int
	for i := 1; i <= until; i++ {
		sum += i
	}
	return sum
}
