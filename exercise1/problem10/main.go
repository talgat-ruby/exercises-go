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

	sum := num1 + num2
	return strconv.Itoa(sum), nil
}

func main() {
	// Примеры использования функции sum
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
