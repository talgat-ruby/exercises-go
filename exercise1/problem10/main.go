package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum("1", "2"))   // "3", nil
	fmt.Println(sum("10", "20")) // "30", nil
	fmt.Println(sum("a", "2"))   // "", "string: a cannot be converted"
}

func sum(x, y string) (string, error) {
	number1, error1 := strconv.Atoi(x)
	if error1 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", x)
	}

	number2, error2 := strconv.Atoi(y)
	if error2 != nil {
		return "", fmt.Errorf("string: %s cannot be converted", y)
	}
	result := number1 + number2
	return strconv.Itoa(result), nil
}
