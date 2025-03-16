package main

import (
	"fmt"
)

func numberSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func main() {
	fmt.Println(numberSquares(2))
	fmt.Println(numberSquares(3))
	fmt.Println(numberSquares(4))
	fmt.Println(numberSquares(5))
}
