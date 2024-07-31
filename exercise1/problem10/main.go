package main

import (
	"fmt"
	"strconv"
)

func sum(str1, str2 string) (string, error) {
	num1, err1 := strconv.Atoi(str1)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", str1)
	}

	num2, err2 := strconv.Atoi(str2)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", str2)
	}

	result := strconv.Itoa(num1 + num2)
	return result, nil
}
