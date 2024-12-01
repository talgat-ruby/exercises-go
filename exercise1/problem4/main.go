package main

import (
	"strings"
)

func detectWord(crowd string) string {
	var result strings.Builder
	for _, char := range crowd {
		if char >= 'a' && char <= 'z' {
			result.WriteRune(char)
		}
	}
	return result.String()
}
