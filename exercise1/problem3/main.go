package main

func numberSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}
