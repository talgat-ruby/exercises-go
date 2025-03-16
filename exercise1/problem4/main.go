package main

import "strings"

func detectWord(random string) string {
	var words []string
	for i := 0; i < len(random); i++ {
		if random[i] >= 97 && random[i] <= 122 {
			words = append(words, random[i:i+1])
		}
	}
	return strings.Join(words, "")
}
