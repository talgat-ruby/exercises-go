package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	totalSquares := numberSquares(n)
	fmt.Println("Total number of squares in a", n, "x", n, "grid:", totalSquares)
}

// countSquares calculates the number of different squares in an n * n square grid.
func numberSquares(n int) int {
	totalSquares := 0
	for size := 1; size <= n; size++ {
		totalSquares += (n - size + 1) * (n - size + 1)
	}
	return totalSquares
}
