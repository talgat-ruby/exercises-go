package main

import (
	"errors"
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	sum1, err := strconv.Atoi(a)
	if err != nil {
		return "", errors.New("string: a cannot be converted")
	}
	sum2, err := strconv.Atoi(b)
	if err != nil {
		return "", errors.New("string: b cannot be converted")
	}
	return fmt.Sprintf("%v", sum1+sum2), nil
}
