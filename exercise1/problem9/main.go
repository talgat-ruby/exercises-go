package main

import (
	"fmt"
)

func toBinary(n int) string {
	return fmt.Sprintf("%08b", n)
}

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
	a := 6
	b := 23

	fmt.Println("a =", toBinary(a))
	fmt.Println("b =", toBinary(b))

	fmt.Println("bitwiseAND =", toBinary(bitwiseAND(a, b)))
	fmt.Println("bitwiseOR =", toBinary(bitwiseOR(a, b)))
	fmt.Println("bitwiseXOR =", toBinary(bitwiseXOR(a, b)))
}
