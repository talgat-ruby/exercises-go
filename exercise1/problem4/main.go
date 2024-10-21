package main

import (
	"unicode"
)

func detectWord(word string) string {
	var result []rune
	for _, ch := range word {
		if unicode.IsLower(ch) {
			result = append(result, ch)
		}
	}
	return string(result)
}
