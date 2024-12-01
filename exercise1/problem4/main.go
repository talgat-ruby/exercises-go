package main

import "unicode"

func detectWord(str string) string {
	var result []rune
	for _, char := range str {
		if unicode.IsLower(char) {
			result = append(result, char)
		}
	}
	return string(result)
}
