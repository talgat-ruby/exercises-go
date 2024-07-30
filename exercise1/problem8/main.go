package main

import (
	"strings"
)

func countVowels(text string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
