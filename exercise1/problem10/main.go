package main

import (
	"errors"
	"strconv"
)

func sum(a string, b string) (string, error) {
	inta, erra := strconv.ParseInt(a, 10, 32)
	intb, errb := strconv.ParseInt(b, 10, 32)
	if erra != nil {
		return "", errors.New("string: a cannot be converted")
	}
	if errb != nil {
		return "", errors.New("string: b cannot be converted")
	}
	result := inta + intb
	return strconv.FormatInt(result, 10), nil
}
