package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	numA, errA := strconv.Atoi(a)
	if errA != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	numB, errB := strconv.Atoi(b)
	if errB != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	sum := numA + numB
	return strconv.Itoa(sum), nil
}

func main() {
	result, err := sum("1", "2")
	fmt.Println(result, err)

	result, err = sum("10", "20")
	fmt.Println(result, err)

	result, err = sum("a", "2")
	fmt.Println(result, err)
}
