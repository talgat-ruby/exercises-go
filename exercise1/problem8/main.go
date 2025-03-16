package main

import "strings"

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0

	// Iterate over each character in the string
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
