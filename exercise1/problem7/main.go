package main

import (
	"slices"
	"strconv"
)

func highestDigit(number int) int {
	digits := []int{}
	numberStr := strconv.Itoa(number)
	for i := 0; i < len(numberStr); i++ {
		d := number % 10

		digits = append(digits, d)
		number /= 10
	}
	slices.Sort(digits)
	return digits[len(digits)-1]
}
