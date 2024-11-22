package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	num1, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	res := num1 + num2
	return strconv.Itoa(res), nil
}
