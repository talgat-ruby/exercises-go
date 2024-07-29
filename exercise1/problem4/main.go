package main

import "unicode"

func detectWord(a string) string {
	sum := ""
	for _, char := range a {
		if unicode.IsLower(char) {
			sum += string(char)
		}
	}
	return sum
}
