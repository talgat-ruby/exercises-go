package main

import (
	"fmt"
	"strconv"
)

// binary function converts a base-10 integer to its base-2 string representation
func binary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	// Примеры использования функции binary
	fmt.Println(binary(1))  // "1"
	fmt.Println(binary(5))  // "101"
	fmt.Println(binary(10)) // "1010"
}
