package main

import "strings"

func countVowels(crowd string) (count int) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

	for _, vowel := range vowels {
		count += strings.Count(crowd, string(vowel))
	}

	return
}
