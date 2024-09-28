package main

func addUp(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum = sum + i
	}
	return sum
}
