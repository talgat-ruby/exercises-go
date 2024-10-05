package main

func numberSquares(number int) int {
	squares := 0
	for i := 1; i <= number; i++ {
		squares += i * i
	}
	return squares
}
