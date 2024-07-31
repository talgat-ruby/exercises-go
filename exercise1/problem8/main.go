package main

import (
	"strings"
)

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0

	for _, char := range s {
		if isVowel(char, vowels) {
			count++
		}
	}

	return count
}

func isVowel(char rune, vowels string) bool {
	return strings.ContainsRune(vowels, char)
}

func main() {
}
