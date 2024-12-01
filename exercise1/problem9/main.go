package main

import (
	"fmt"
)

// bitwiseAND выполняет побитовое И (AND) для двух чисел
func bitwiseAND(a, b int) int {
	return a & b
}

// bitwiseOR выполняет побитовое ИЛИ (OR) для двух чисел
func bitwiseOR(a, b int) int {
	return a | b
}

// bitwiseXOR выполняет побитовое исключающее ИЛИ (XOR) для двух чисел
func bitwiseXOR(a, b int) int {
	return a ^ b
}

func main() {
	a := 6
	b := 23

	fmt.Printf("bitwiseAND(%d, %d) = %08b\n", a, b, bitwiseAND(a, b))
	fmt.Printf("bitwiseOR(%d, %d) = %08b\n", a, b, bitwiseOR(a, b))
	fmt.Printf("bitwiseXOR(%d, %d) = %08b\n", a, b, bitwiseXOR(a, b))
}
