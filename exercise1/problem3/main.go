package main

import "fmt"

func numberSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}

func main() {
	fmt.Println(numberSquares(2))
	fmt.Println(numberSquares(3))
	fmt.Println(numberSquares(4))
	fmt.Println(numberSquares(5))
}
