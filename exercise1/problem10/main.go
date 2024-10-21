package main

import "strconv"

func sum(s1, s2 string) (string, error) {
	n1, err := strconv.Atoi(s1)
	if err != nil {
		return "", err
	}

	n2, err := strconv.Atoi(s2)
	if err != nil {
		return "", err
	}

	n3 := n1 + n2
	return strconv.Itoa(n3), nil
}
