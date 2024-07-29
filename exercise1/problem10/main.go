package main

import (
	"fmt"
	"strconv"
)

func sum(a,b string) (string,error) {
	num1, err1 := strconv.Atoi(a)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	num2, err2 := strconv.Atoi(b)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	result := num1 + num2
	return strconv.Itoa(result), nil
}
