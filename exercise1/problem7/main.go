package main

import (
	"strconv"
)

func highestDigit(num int) int {
	numToStr := strconv.Itoa(num)
	highest := 0

	for _, val := range numToStr {
		digit, _ := strconv.Atoi(string(val))
		if digit > highest {
			highest = digit
		}
	}
	return highest
}
