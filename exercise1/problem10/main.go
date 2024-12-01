package main

import (
	"fmt"
	"strconv"
)

// sum function adds two numbers represented as strings
// and returns the result as a string or an error if conversion fails
func sum(a, b string) (string, error) {
	// Convert the first string to an integer
	num1, err1 := strconv.Atoi(a)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	// Convert the second string to an integer
	num2, err2 := strconv.Atoi(b)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	// Calculate the sum and return as a string
	result := num1 + num2
	return strconv.Itoa(result), nil
}

func main() {
	// Examples of using the sum function
	result, err := sum("1", "2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result) // "3"
	}

	result, err = sum("10", "20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result) // "30"
	}

	result, err = sum("a", "2")
	if err != nil {
		fmt.Println(err) // "string: a cannot be converted"
	} else {
		fmt.Println(result)
	}
}
