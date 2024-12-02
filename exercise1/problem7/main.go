package main

import "strconv"

func highestDigit(num int) int {
	numStr := strconv.Itoa(num)
	maxNum := 0
	for i := 0; i < len(numStr); i++ {
		digit, _ := strconv.Atoi(string(numStr[i]))
		if maxNum < digit {
			maxNum = digit
		}
	}
	return maxNum
}
