package main

func numberSquares(n int) int {
	if n == 1 {
		return 1
	} else {
		return n*n + numberSquares(n-1)
	}
}
