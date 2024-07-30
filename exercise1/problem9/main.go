package main

import (
	"fmt"
)

// bitwiseAND function returns the bitwise AND of two integers
func bitwiseAND(a, b int) int {
	return a & b
}

// bitwiseOR function returns the bitwise OR of two integers
func bitwiseOR(a, b int) int {
	return a | b
}

// bitwiseXOR function returns the bitwise XOR of two integers
func bitwiseXOR(a, b int) int {
	return a ^ b
}

func main() {
	// Examples of using the bitwise functions
	fmt.Printf("bitwiseAND(6, 23) = %08b\n", bitwiseAND(6, 23)) // 00000110
	fmt.Printf("bitwiseOR(6, 23) = %08b\n", bitwiseOR(6, 23))   // 00010111
	fmt.Printf("bitwiseXOR(6, 23) = %08b\n", bitwiseXOR(6, 23)) // 00010001
}
