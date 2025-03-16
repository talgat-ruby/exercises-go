package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, string) {
	intA, errA := strconv.Atoi(a)
	if errA != nil {
		return "", fmt.Sprintf("string: %s cannot be converted", a)
	}
	intB, errB := strconv.Atoi(b)
	if errB != nil {
		return "", fmt.Sprintf("string: %s cannot be converted", b)
	}
	sum := intA + intB
	return strconv.Itoa(sum), ""
}
