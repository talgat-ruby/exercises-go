package main

import (
	"unicode"
)

func detectWord(word string) string {
	result := ""
	for _, ch := range word {
		if unicode.IsLower(ch) {
			result += string(ch)
		}
	}
	return result
}

func main() {
}
