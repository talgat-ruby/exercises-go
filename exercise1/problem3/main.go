package main

func numberSquares(n int) int {
	total := 0
	for k := 1; k <= n; k++ {
		total += (n - k + 1) * (n - k + 1)
	}
	return total
}
