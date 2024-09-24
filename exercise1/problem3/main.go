package main

func numberSquares(n int) int {
	squares := 0
	for i := 1; i <= n; i++ {
		squares += (n - i + 1) * (n - i + 1)
	}
	return squares
}
