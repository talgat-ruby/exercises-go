package main

import "strconv"

func sum(a, b string) (string, error) {

	n1, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}
	n2, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}
	res:=n1+n2
	return strconv.Itoa(res),nil
}

