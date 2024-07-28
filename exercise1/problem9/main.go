package main

import (
	"fmt"
)

func main() {
	a := 6
	b := 23
	fmt.Println(bitwiseAND(a, b))
	fmt.Println(bitwiseOR(a, b))
	fmt.Println(bitwiseXOR(a, b))
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
