package main

import (
	"unicode"
)

func detectWord(input string) string {
	var result string

	for _, char := range input {
		if unicode.IsLower(char) {
			result += string(char)
		}
	}

	return result
}
