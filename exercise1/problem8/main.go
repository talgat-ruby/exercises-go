package main

import "strings"

func countVowels(a string) int {
	vowels := "aeiouAEIOU"
	sum := 0
	for _, ch := range a {
		if strings.ContainsRune(vowels, ch) {
			sum++
		}
	}
	return sum
}
