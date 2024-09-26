package main

import "strconv"

func highestDigit(n int) int {
	str := strconv.Itoa(n)
	maxDigit := 0

	for _, ch := range str {
		digit, _ := strconv.Atoi(string(ch))
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}
