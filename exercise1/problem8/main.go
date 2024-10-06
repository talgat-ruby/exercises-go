package main

import "strings"

func countVowels(str string) int {
	var c int = 0
	var vowels = []string{"a", "e", "i", "o", "u"}
	for _, ch := range vowels {
		c += strings.Count(str, ch)
	}
	return c
}
