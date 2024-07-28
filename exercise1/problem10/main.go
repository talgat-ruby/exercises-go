package main

import "strconv"

func sum(a string, b string) (string, error) {
	aI, err := strconv.Atoi(a)

	if err != nil {
		return "", err
	}

	bI, err := strconv.Atoi(b)

	if err != nil {
		return "", err
	}

	return strconv.Itoa(aI + bI), nil
}
