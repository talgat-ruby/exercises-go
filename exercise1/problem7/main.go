package main

import (
	"fmt"
	"strconv"
)

func highestDigit(num int) int {
	str := strconv.Itoa(num)
	maxDigit := 0

	for _, char := range str {
		digit, _ := strconv.Atoi(string(char))
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
