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
