package main

import "unicode"

func detectWord(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLower(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
