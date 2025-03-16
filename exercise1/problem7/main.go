package main

import "strconv"

func highestDigit(number int) int {
	strNum := strconv.Itoa(number)
	highDigit := 0

	for _, digitRune := range strNum {
		digit, _ := strconv.Atoi(string(digitRune))
		if digit > highDigit {
			highDigit = digit
		}
	}

	return highDigit
}
