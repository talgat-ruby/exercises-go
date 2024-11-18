package main

import "fmt"

func main() {
	fmt.Println(numberSquares(2))
}

func numberSquares(n int) int {
	// 1^2 + 2^2 + 3^2 + … + N^2
	if n == 0 {
		return 0
	}
	var total = 0
	for i := 1; i <= n; i++ {
		total += i * i
	}
	return total
}
