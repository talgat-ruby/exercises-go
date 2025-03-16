package main

import "math"

func numberSquares(n int) int {
	squares := 0
	for i := 1; i <= n; i++ {
		squares += int(math.Pow(float64(i), 2))
	}
	return squares
}
