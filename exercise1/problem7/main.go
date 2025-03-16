package main

import "strconv"

func highestDigit(n int) int {
	numStr := strconv.Itoa(n)

	var digits []int

	for _, char := range numStr {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		digits = append(digits, digit)
	}

	maxDigit := 0

	for _, digit := range digits {
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}
