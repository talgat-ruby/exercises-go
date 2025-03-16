package main

import (
	"fmt"
	"strconv"
)

func sum(x string, y string) (string, error) {
	xInt, err := strconv.Atoi(x)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", x)
	}
	yInt, err := strconv.Atoi(y)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", y)
	}
	res := strconv.Itoa(xInt + yInt)
	return res, nil
}
