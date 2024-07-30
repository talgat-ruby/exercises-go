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
	fmt.Printf("%08b\n", bitwiseAND(6, 23))
	fmt.Printf("%08b\n", bitwiseOR(6, 23))
	fmt.Printf("%08b\n", bitwiseXOR(6, 23))
}
