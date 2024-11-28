package main

import "unicode"

func detectWord(s string) string {
	word := ""
	for _, char := range s {
		if unicode.IsLower(char) {
			word += string(char)
		}
	}
	return word
}
