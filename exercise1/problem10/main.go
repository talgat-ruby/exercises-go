package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	num_a, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: " + a + "cannot be converted")
	}
	num_b, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: " + b + "cannot be converted")
	}
	return strconv.Itoa(num_b + num_a), nil
}
