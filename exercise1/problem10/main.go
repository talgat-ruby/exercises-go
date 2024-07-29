package main

import (
	"errors"
	"strconv"
)

func sum(a, b string) (string, error) {
	intA, err := strconv.Atoi(a)
	if err != nil {
		return "", errors.New("string: " + a + " cannot be converted")
	}
	intB, err := strconv.Atoi(b)
	if err != nil {
		return "", errors.New("string: " + b + " cannot be converted")
	}
	return strconv.Itoa(intA + intB), nil
}
