package main

import "strconv"

func sum(a, b string) (string, error) {
	numA, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}

	numB, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}

	res := numA + numB
	str := strconv.Itoa(res) 
	return str, nil
}
