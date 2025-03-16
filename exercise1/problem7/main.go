package main

import "strconv"

func highestDigit(num int) int {
	temp := strconv.Itoa(num)
	high := "0"
	for i := range temp {
		if string(temp[i]) > high {
			high = string(temp[i])
		}
	}
	res, _ := strconv.Atoi(high)
	return res
}
