package main

func numberSquares(n int) int {
	quantity := 0
	for i := 1; i <= n; i++ {
		quantity += (n - i + 1) * (n - i + 1)
	}
	return quantity
}
