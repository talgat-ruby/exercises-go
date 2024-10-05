package main

import (
	"unicode"
)

func detectWord(word string) (detectWord string) {
	for _, char := range word {
		if unicode.IsLower(char) {
			detectWord += string(char)
		}
	}
	return
}
