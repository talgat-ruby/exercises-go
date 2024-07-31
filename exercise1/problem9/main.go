package main

import "fmt"

func bitwiseAND(a, b int) string {
	res := a & b
	return fmt.Sprintf("%08b", res)
}

func bitwiseOR(a, b int) string {
	res := a | b
	return fmt.Sprintf("%08b", res)
}

func bitwiseXOR(a, b int) string {
	res := a ^ b
	return fmt.Sprintf("%08b", res)
}
