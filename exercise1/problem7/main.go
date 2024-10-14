package main

import (
	"fmt"
	"strconv"
)

func highestDigit(num int) int {
	highest := 0
	strnum := fmt.Sprintf("%d", num)
	for _, val := range strnum {
		newVal, _ := strconv.Atoi(string(val))
		if newVal > highest {
			highest = int(newVal)
		}
	}
	return highest
}
