package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	intA, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	intB, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	result := intA + intB

	return strconv.Itoa(result), nil
}
