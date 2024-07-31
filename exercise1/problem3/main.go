package main

import "fmt"

func numberSquares(number int) int {
	total := 0

	for i := 1; i <= number; i++ {
		total += i * i
	}

	return total
}

func main() {
	fmt.Println(numberSquares(2))
}
