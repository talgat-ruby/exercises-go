package main

import (
	"fmt"
	"strconv"
)

func sum(a string, b string) (string, error) {

	aNum, err := strconv.Atoi(a)
	if err != nil {

		return "", fmt.Errorf("string: %s cannot be converted", a)
	}

	bNum, err := strconv.Atoi(b)
	if err != nil {

		return "", fmt.Errorf("string: %s cannot be converted", b)
	}

	return strconv.Itoa(aNum + bNum), nil
}
