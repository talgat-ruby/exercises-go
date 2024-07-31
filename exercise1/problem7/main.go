package main

import (
	"fmt"
	"strconv"
)

func highestDigit(n int) int {
	str := strconv.Itoa(n)
	maxDigit := 0
	for _, char := range str {
		digit := int(char - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}

func main() {
	fmt.Println(highestDigit(100))
}
