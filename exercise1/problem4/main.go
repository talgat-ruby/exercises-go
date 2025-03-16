package main

import "unicode"

func detectWord(crowd string) string {
	word := ""
	for _, char := range crowd {
		if unicode.IsLower(char) {
			word += string(char)
		}
	}
	return word
}
