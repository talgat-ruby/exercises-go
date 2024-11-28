package main

import (
	"fmt"
	"strconv"
)

func sum(str1, str2 string) (int, error) {
	num1, err := strconv.Atoi(str1)
	if err != nil {
		return 0, fmt.Errorf("string: %s cannot be converted", str1)
	}

	num2, err := strconv.Atoi(str2)
	if err != nil {
		return 0, fmt.Errorf("string: %s cannot be converted", str2)
	}

	return num1 + num2, nil
}
