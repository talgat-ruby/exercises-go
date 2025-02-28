package main

import (
	"fmt"
)

func numberSquares(n int) int {
	sum := 0
	for k := 1; k <= n; k++ {
		sum += (n - k + 1) * (n - k + 1)
	}
	return sum
}

func main() {
	fmt.Println(numberSquares(2))
	fmt.Println(numberSquares(3))
	fmt.Println(numberSquares(4))
	fmt.Println(numberSquares(5))
}
