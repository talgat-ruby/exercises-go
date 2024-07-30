package main

import "strconv"

func sum(a, b string) (string, error) {
	aInt, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(aInt + bInt), nil
}
