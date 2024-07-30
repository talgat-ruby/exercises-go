package main

import "strconv"

func highestDigit(n int) int {
	numStr := strconv.Itoa(n)

	maxDigit := 0
	for _, char := range numStr {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}
