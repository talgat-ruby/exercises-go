package main

import "strings"

func countVowels(s string) int {
	count := 0
	vowels := "aeiou"
	for _, c := range s {
		if strings.Contains(vowels, string(c)) {
			count++
		}
	}
	return count
}
