package main

import (
	"fmt"
	"strconv"
)

// highestDigit function returns the highest digit in a given number
func highestDigit(n int) int {
	// Convert the number to a string
	numStr := strconv.Itoa(n)

	// Initialize the highest digit
	maxDigit := 0

	// Iterate through each character in the string
	for _, char := range numStr {
		// Convert the character to an integer digit
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			// Handle conversion error
			continue
		}
		// Update the highest digit if the current digit is greater
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}

func main() {
	// Examples of using the highestDigit function
	fmt.Println(highestDigit(379))    // 9
	fmt.Println(highestDigit(2))      // 2
	fmt.Println(highestDigit(377401)) // 7
}
