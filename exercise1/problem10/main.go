package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	ai, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	result := ai + bi
	return strconv.Itoa(result), nil
}
