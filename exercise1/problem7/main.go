package main

import "slices"

func highestDigit(num int) int {
	if num != 0 {
		numS := ItoSlice(num)
		return slices.Max(numS)
	}
	return 0
}

func ItoSlice(n int) []int {
	var numS []int
	for n >= 1 {
		i := n % 10
		numS = append(numS, i)
		n = n / 10
	}
	return numS
}
