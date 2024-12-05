package main

import "unicode"

func detectWord(str string) string {
	var result string
	for _, ch := range str {
		if unicode.IsLower(ch) {
			result += string(ch)
		}
	}
	return result
}
