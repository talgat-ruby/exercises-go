package main

import "math"

func numberSquares(n int) int {
	s := 0

	for i := 1; i <= n; i++ {
		s += int(math.Pow(float64(i), 2))
	}
	return int(s)
}
