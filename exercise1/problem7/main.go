package main

import "strconv"

func highestDigit(number int) int {
	numStr := strconv.Itoa(number)

	highest := 0

	for _, char := range numStr {
		digit := int(char - '0')
		if digit > highest {
			highest = digit
		}
	}

	return highest
}
