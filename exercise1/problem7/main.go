package main

import (
	"fmt"
	"strconv"
)

func highestDigit(num int) int {
	numStr := strconv.Itoa(num)
	maxDigit := 0

	for _, ch := range numStr {
		digit, _ := strconv.Atoi(string(ch))
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}

func main() {
	fmt.Println(highestDigit(379))
	fmt.Println(highestDigit(2))
	fmt.Println(highestDigit(377401))
}
