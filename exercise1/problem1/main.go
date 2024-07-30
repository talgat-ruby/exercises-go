package main

func addUp(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i
	}
	return sum
}
