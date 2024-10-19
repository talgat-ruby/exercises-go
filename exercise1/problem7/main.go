package main

import (
	"fmt"
	"strconv"
)

func highestDigit(num int) int {
	numStr := strconv.Itoa(num)

	highest := 0

	for _, char := range numStr {
		digit, _ := strconv.Atoi(string(char))

		if digit > highest {
			highest = digit
		}
	}

	return highest
}

func main() {
	fmt.Println(highestDigit(379))
	fmt.Println(highestDigit(2))
	fmt.Println(highestDigit(377401))
}
