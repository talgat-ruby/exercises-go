package main

import "strings"

//func detectWord() {}

func detectWord(s string) string {
	// Step 1: Extract lowercase letters from the crowd
	var a strings.Builder
	for _, char := range s {
		if string(char) == strings.ToLower(string(char)) {
			a.WriteString(string(char))
		}
	}
	return a.String()
}
