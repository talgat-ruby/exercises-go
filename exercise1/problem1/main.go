package main

func addUp(num int) int {
	n, sum := 1, 0
	for n <= num {
		sum += n
	}
	return sum
}
