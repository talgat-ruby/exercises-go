package main

import "strconv"

func sum(str1 string, str2 string) (string, error) {
	number1, err := strconv.Atoi(str1)
	if err != nil {
		return "", err
	}

	number2, err := strconv.Atoi(str2)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(number1 + number2), nil
}
