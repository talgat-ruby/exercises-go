package main

import "strconv"

func highestDigit(x int) int {
	ans := '0'
	s := strconv.Itoa(x)
	for _, c := range s {
		if c > ans {
			ans = c
		}
	}
	maxDigitInt, _ := strconv.Atoi(string(ans))

	return maxDigitInt
}
