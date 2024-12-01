package main

import (
	"fmt"
	"strconv"
)

func highestDigit(n int) int {
	numStr := strconv.Itoa(n)
	highest := 0

	for _, char := range numStr {
		digit := int(char - '0')
		if digit > highest {
			highest = digit
		}
	}

	return highest
}

func main() {
	fmt.Println(highestDigit(379))    // 9
	fmt.Println(highestDigit(2))      // 2
	fmt.Println(highestDigit(377401)) // 7
}
