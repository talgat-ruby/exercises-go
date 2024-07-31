package main

import (
	"fmt"
	"strconv"
)

func sum(str1 string, str2 string) (string, error) {
	num1, err1 := strconv.Atoi(str1)
	if err1 != nil {
		return "", fmt.Errorf("string: a cannot be converted")
	}
	num2, err2 := strconv.Atoi(str2)
	if err2 != nil {
		return "", fmt.Errorf("string: a cannot be converted")
	}
	summ := num1 + num2
	result := strconv.Itoa(summ)
	return result, nil
}
