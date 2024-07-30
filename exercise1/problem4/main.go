package main

import (
	"strings"
	"unicode"
)

func detectWord(str string) string {
	var result strings.Builder
	for _, char := range str {
		if unicode.IsLower(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}
