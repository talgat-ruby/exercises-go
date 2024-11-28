package main

import (
	"fmt"
	"strconv"
)

// sum takes two strings, converts them to integers, sums them, and returns the result as a string.
// If either string cannot be converted, it returns an error message.
func sum(a, b string) (string, error) {
	num1, err1 := strconv.Atoi(a)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	num2, err2 := strconv.Atoi(b)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	// Sum the numbers and return the result as a string
	result := strconv.Itoa(num1 + num2)
	return result, nil
}

func main() {
	// Test cases
	fmt.Println(sum("1", "2"))    // Output: "3", nil
	fmt.Println(sum("10", "20"))  // Output: "30", nil
	fmt.Println(sum("a", "2"))    // Output: "", "string: a cannot be converted"
	fmt.Println(sum("15", "b"))   // Output: "", "string: b cannot be converted"
	fmt.Println(sum("5", "10abc")) // Output: "", "string: 10abc cannot be converted"
}
