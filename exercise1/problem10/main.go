package main

import (
	"fmt"
	"strconv"
)

func sum(s1, s2 string) (string, error) {
	n1, err := strconv.Atoi(s1)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", s1)
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", s2)
	}

	return strconv.Itoa(n1 + n2), nil
}
