package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(highestDigit(379))
	fmt.Println(highestDigit(2))
	fmt.Println(highestDigit(377401))
}

func highestDigit(number int) int {
	numStr := strconv.Itoa(number)
	maxDigit := 0

	for _, digit := range numStr {
		digit, _ := strconv.Atoi(string(digit))
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}
