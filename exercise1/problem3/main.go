package main

import (
	"fmt"
)

// numberSquares function calculates the number of different squares in an n * n square grid
func numberSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func main() {
	// Примеры использования функции numberSquares
	fmt.Println(numberSquares(2)) // 5
	fmt.Println(numberSquares(3)) // 14
	fmt.Println(numberSquares(4)) // 30
	fmt.Println(numberSquares(5)) // 55
}
