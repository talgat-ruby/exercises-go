package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 394587 // Example number
	highest := highestDigit(number)
	fmt.Println("The highest digit in the number is:", highest)
}

// highestDigit takes a number as an argument and returns the highest digit in that number.
func highestDigit(number int) int {
	highest := 0

	// Convert the number to a string
	numberStr := strconv.Itoa(number)

	// Iterate through each character in the string
	for _, char := range numberStr {
		// Convert the character back to an integer
		digit := int(char - '0')

		// Update the highest digit found
		if digit > highest {
			highest = digit
		}
	}

	return highest
}
