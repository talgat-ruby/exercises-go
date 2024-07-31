package main

import (
	"errors"
	"fmt"
	"strconv"
)

func sum(str1, str2 string) (string, error) {
	num1, err1 := strconv.Atoi(str1)
	if err1 != nil {
		return "", errors.New(fmt.Sprintf("%s cannot be converted", str1))
	}

	num2, err2 := strconv.Atoi(str2)
	if err2 != nil {
		return "", errors.New(fmt.Sprintf("%s cannot be converted", str2))
	}

	result := num1 + num2
	return strconv.Itoa(result), nil
}

func main() {
	fmt.Println(sum("1", "2"))
}
