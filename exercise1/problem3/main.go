package main

func numberSquares(number int) int {
	var total int
	for i := 1; i <= number; i++ {
		total += i * i
	}
	return total
}
