package main

import (
	"fmt"
	"strconv"
)

func sum(m, n string) (string, error) {
	intM, err := strconv.Atoi(m)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", m)
	}

	intN, err := strconv.Atoi(n)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", n)
	}

	result := intM + intN

	return strconv.Itoa(result), nil
}
