package main

import (
	"math"
	"strconv"
)

func highestDigit(num int) int {
	num = int(math.Abs(float64(num)))

	numStr := strconv.Itoa(num)

	highest := 0
	for _, char := range numStr {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		if digit > highest {
			highest = digit
		}
	}

	return highest
}
