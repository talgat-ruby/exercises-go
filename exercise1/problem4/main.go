package main

import (
	"unicode"
)

func detectWord(crowd string) string {
	var word string
	for _, letter := range crowd {
		if unicode.IsLower(letter) {
			word += string(letter)
		}
	}
	return word
}
