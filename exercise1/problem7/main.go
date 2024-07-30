package main

import (
	"fmt"
	"strconv"
)

func highestDigit(num int) int {

	numStr := strconv.Itoa(num)
	maxDigit := 0

	for _, char := range numStr {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting character to digit:", err)
			return -1
		}
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}
