package main

import (
	"fmt"
)

func bitwiseAND(a, b int) int {
	return a & b
}

func bitwiseOR(a, b int) int {
	return a | b
}

func bitwiseXOR(a, b int) int {
	return a ^ b
}

func main() {
	a, b := 6, 23

	fmt.Printf("bitwiseAND(%d, %d) = %d\n", a, b, bitwiseAND(a, b))
	fmt.Printf("bitwiseOR(%d, %d) = %d\n", a, b, bitwiseOR(a, b))
	fmt.Printf("bitwiseXOR(%d, %d) = %d\n", a, b, bitwiseXOR(a, b))
}
