package main

import "unicode"

func countVowels(word string) int {

	cnt := 0

	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	for _, ch := range word {
		ch = unicode.ToLower(ch)

		_, ok := vowels[string(ch)]
		if ok {
			cnt += 1
		}
	}

	return cnt
}
