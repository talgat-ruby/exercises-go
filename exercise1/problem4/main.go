package main

import (
	"unicode"
)

func detectWord(crowd string) string {
	var result string
	for _, s := range crowd {
		if !unicode.IsUpper(s) {
			result = result + string(s)
		}
	}
	return result
}
