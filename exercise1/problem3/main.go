package main

func numberSquares(n int) int {
	totalSquares := 0
	for k := 1; k <= n; k++ {
		totalSquares += k * k
	}
	return totalSquares
}
