package main

func numberSquares(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += i * i
	}
	return sum
}
