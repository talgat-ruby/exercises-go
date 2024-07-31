package main

import "strings"

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	return count
}
