package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	intA, err1 := strconv.Atoi(a)
	if err1 != nil {
		return "", err(a)
	}

	intB, err2 := strconv.Atoi(b)
	if err2 != nil {
		return "", err(b)
	}

	return strconv.Itoa(intA + intB), nil
}

func err(s string) error {
	return fmt.Errorf("string: %s cannot be converted", s)
}
