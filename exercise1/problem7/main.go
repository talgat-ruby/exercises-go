package main

import (
	"fmt"
	"strconv"
)

func highestDigit(num int) int {
	numStr := strconv.Itoa(num)
	maxDigit := '0'
	for _, char := range numStr {
		if char > maxDigit {
			maxDigit = char
		}
	}
	return int(maxDigit - '0')
}

func main() {
	fmt.Println(highestDigit(379))
	fmt.Println(highestDigit(2))
	fmt.Println(highestDigit(377401))
}
