package main

func numberSquares(n int) int {
	totalSquares := 0
	for size := 1; size <= n; size++ {
		totalSquares += (n - size + 1) * (n - size + 1)
	}
	return totalSquares
}
