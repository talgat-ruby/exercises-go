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
	fmt.Printf("bitwiseAND(6, 23) = %d\n", bitwiseAND(6, 23))
	fmt.Printf("bitwiseOR(6, 23) = %d\n", bitwiseOR(6, 23))
	fmt.Printf("bitwiseXOR(6, 23) = %d\n", bitwiseXOR(6, 23))
}
