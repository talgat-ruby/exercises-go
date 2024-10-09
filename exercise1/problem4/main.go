package main

import "unicode"

func detectWord(text string) string {
	result := ""
	for _, char := range text {
		if unicode.IsLower(char) {
			result += string(char)
		}
	}
	return result
}
