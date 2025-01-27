package main

import "strings"

func potatoes(text string) int {
	way := 2
	var res int
	switch way {
	case 1:
		res = strings.Count(text, "potato")
	case 2:
		res = way2(text)
	}

	return res
}

func way2(text string) int {
	substr := "potato"
	substrLen := len(substr)
	count := 0

	for i := 0; i <= len(text)-substrLen; i++ {
		if text[i:i+substrLen] == substr {
			count++
		}
	}
	return count
}
