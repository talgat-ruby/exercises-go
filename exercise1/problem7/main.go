package main

import (
	"strconv"
)

func highestDigit(number int) int {
	numStr := strconv.Itoa(number)
	highest := 0

	for _, digit := range numStr {
		digitInt, _ := strconv.Atoi(string(digit))
		if digitInt > highest {
			highest = digitInt
		}
	}

	return highest
}
