package main

import (
	"fmt"
	"strings"
)

// countVowels takes a string and returns the count of vowels in it
func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0

	// Loop through each character in the string
	for _, char := range s {
		// Check if the character is in the vowels string
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countVowels("Celebration")) // Output: 5
	fmt.Println(countVowels("Palm"))        // Output: 1
	fmt.Println(countVowels("Prediction"))  // Output: 4
}
