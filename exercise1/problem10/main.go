package main

import (
	"errors"
	"strconv"
)

func sum(a string, b string) (string, error) {
	sA, errA := strconv.Atoi(a)
	sB, errB := strconv.Atoi(b)
	if errA != nil || errB != nil {
		return "", errors.Join(errA, errB)
	}
	return strconv.Itoa(sA + sB), nil
}
