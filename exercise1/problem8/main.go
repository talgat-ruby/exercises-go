package main

import "strings"

func countVowels(input string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
