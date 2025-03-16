package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!" // Example string
	vowelCount := countVowels(text)
	fmt.Println("Number of vowels in the string:", vowelCount)
}

// countVowels takes a string and returns the number of vowels contained within it.
func countVowels(text string) int {
	vowels := "aeiouAEIOU"
	count := 0

	// Iterate through each character in the string
	for _, char := range text {
		// Check if the character is a vowel
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
