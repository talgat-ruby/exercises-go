package main

import (
	"strconv"
)

func highestDigit(number int) int {
	strnum := strconv.Itoa(number)
	maxnum := int(strnum[0] - '0')

	for i := 1; i < len(strnum); i++ {
		digit := int(strnum[i] - '0')
		if digit > maxnum {
			maxnum = digit
		}
	}
	return maxnum
}
