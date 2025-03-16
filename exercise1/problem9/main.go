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
	fmt.Printf("bitwiseAND(6, 23) = %08b\n", bitwiseAND(6, 23)) // 00000110
	fmt.Printf("bitwiseOR(6, 23) = %08b\n", bitwiseOR(6, 23))   // 00010111
	fmt.Printf("bitwiseXOR(6, 23) = %08b\n", bitwiseXOR(6, 23)) // 00010001
}
