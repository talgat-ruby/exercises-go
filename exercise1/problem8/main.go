package main

import "strings"

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0

	for _, char := range strings.ToLower(s) {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
