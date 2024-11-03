package main

import (
	"unicode"
)

func detectWord(crowd string) string {
	var result string
	for _, char := range crowd {
		if unicode.IsLower(char) {
			result += string(char)
		}
	}
	return result
}
