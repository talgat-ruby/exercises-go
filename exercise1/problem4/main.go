package main

import (
	"unicode"
)

func detectWord(str string) string {
	var word string
	for i := 0; i < len(str); i++ {
		if unicode.IsLower(rune(str[i])) {
			word = word + string(str[i])
		}
	}
	return word
}
