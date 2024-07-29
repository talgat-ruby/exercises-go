package main

import "strconv"

func highestDigit(in int) int {
	digits := []rune(strconv.Itoa(in))
	highest := 0
	for _, d := range digits {
		digit, _ := strconv.Atoi(string(d))
		if digit > highest {
			highest = digit
		}
	}
	return highest
}
