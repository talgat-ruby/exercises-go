package main

import (
	"strconv"
)

func highestDigit(num int) int {
	numStr := strconv.Itoa(num)
	maxDigit := 0

	for _, digit := range numStr {
		digitValue := int(digit - '0')
		if digitValue > maxDigit {
			maxDigit = digitValue
		}
	}

	return maxDigit
}
