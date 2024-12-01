package main

import "strconv"

func binary(number int) string {
	if number == 0 {
		return "0"
	}
	binaryStr := ""
	for number > 0 {
		binaryStr = strconv.Itoa(number%2) + binaryStr
		number = number / 2
	}
	return binaryStr
}
