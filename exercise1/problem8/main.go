package main

import (
	"strings"
)

func countVowels(str string) int {
	str = strings.ToLower(str)
	vowels := "aoiue"
	count := 0

	for _, ch := range str {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	return count
}
