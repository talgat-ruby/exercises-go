package main

import (
	"fmt"
	"strconv"
)

func binary(a string) (string, error) {
	num, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("invalid decimal number: %s", a)
	}

	binary := strconv.FormatInt(int64(num), 2)
	return binary, nil
}
