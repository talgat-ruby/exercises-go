package main

import (
	"strings"
)

func countVowels(text string) int {
	vowels := "aeiou"

	count := 0

	text = strings.ToLower(text)

	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
