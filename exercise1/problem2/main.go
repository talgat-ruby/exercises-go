package main

import "strconv"

func binary(decNumber int) string {
	if decNumber == 0 {
		return "0"
	}
	binaryNumber := ""
	for decNumber > 0 {
		binaryNumber = strconv.Itoa(decNumber%2) + binaryNumber
		decNumber = decNumber / 2
	}

	return binaryNumber

}
