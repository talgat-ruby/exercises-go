package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	num1, err1 := strconv.Atoi(a)
	if err1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	num2, err2 := strconv.Atoi(b)
	if err2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	sum := num1 + num2
	result := strconv.Itoa(sum)
	return result, nil
}

func main() {

	res, err := sum("1", "2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = sum("10", "20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res, err = sum("a", "2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
