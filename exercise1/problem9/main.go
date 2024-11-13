package main

import (
	"fmt"
)

// bitwiseAND performs a bitwise AND operation on two integers
func bitwiseAND(a, b int) int {
	return a & b
}

// bitwiseOR performs a bitwise OR operation on two integers
func bitwiseOR(a, b int) int {
	return a | b
}

// bitwiseXOR performs a bitwise XOR operation on two integers
func bitwiseXOR(a, b int) int {
	return a ^ b
}

// toBinaryString converts an integer to its binary representation as a string, padded to 8 bits
func toBinaryString(n int) string {
	return fmt.Sprintf("%08b", n)
}

func main() {
	a, b := 6, 23
	
	fmt.Printf("bitwiseAND(%d, %d) = %s\n", a,
