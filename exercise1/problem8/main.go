package main

import (
	"fmt"
	"strings"
)

// countVowels function returns the number of vowels in the given string
func countVowels(s string) int {
	// Define a set of vowels (both uppercase and lowercase)
	vowels := "aeiouAEIOU"
	count := 0

	// Iterate through each character in the string
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func main() {
	// Examples of using the countVowels function
	fmt.Println(countVowels("Celebration")) // 5
	fmt.Println(countVowels("Palm"))        // 1
	fmt.Println(countVowels("Prediction"))  // 4
}
