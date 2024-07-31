package main

import (
	"fmt"
	"strconv"
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
	fmt.Println(strconv.FormatInt(int64(bitwiseAND(6, 23)), 2))
	fmt.Println(strconv.FormatInt(int64(bitwiseOR(6, 23)), 2))
	fmt.Println(strconv.FormatInt(int64(bitwiseXOR(6, 23)), 2))
}
