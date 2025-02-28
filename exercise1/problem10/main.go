package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	// Convert the first string to an integer
	intA, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	// Convert the second string to an integer
	intB, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	// Sum the integers
	result := intA + intB

	// Convert the result back to a string and return it
	return strconv.Itoa(result), nil
}
