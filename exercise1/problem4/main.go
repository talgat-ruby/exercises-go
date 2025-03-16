package main

import "unicode"

func detectWord(str string) string {
	var res []rune

	for _, char := range str {
		if unicode.IsLower(char) {
			res = append(res, char)
		}
	}

	return string(res)
}
