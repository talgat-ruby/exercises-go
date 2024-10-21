package main

import (
	"fmt"
	"strconv"
)

func sum(s1, s2 string) (string, error) {
	num1, err1 := strconv.Atoi(s1)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", s1)
	}

	num2, err2 := strconv.Atoi(s2)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", s2)
	}

	result := num1 + num2

	return strconv.Itoa(result), nil
}
