package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)

	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	return strconv.Itoa(num1 + num2), nil
}

func main() {
	result, err := sum("1", "2")
	fmt.Println(result, err)

	result, err = sum("10", "20")
	fmt.Println(result, err)

	result, err = sum("a", "2")
	fmt.Println(result, err)
}
