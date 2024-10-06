package main

import "fmt"

func main() {
	fmt.Println(numberSquares(2))
	fmt.Println(numberSquares(3))
	fmt.Println(numberSquares(4))
	fmt.Println(numberSquares(5))
}

func numberSquares(n int) int {
	totalSquares := 0
	for i := 1; i <= n; i++ {
		totalSquares += (n - i + 1) * (n - i + 1)
	}
	return totalSquares
}
