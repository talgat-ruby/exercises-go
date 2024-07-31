package main

import "strconv"

func highestDigit(number int) int {
	maxDigit := 0
	numStr := strconv.Itoa(number)

	for _, digit := range numStr {
		digitInt := int(digit - '0') // Convert rune to int
		if digitInt > maxDigit {
			maxDigit = digitInt
		}
	}

	return maxDigit
}
