package main

import "strings"

func potatoes(word string) int {

	// return strings.Count(word, "potato")

	substr := "potato"
	count := 0

	for {
		index := strings.Index(word, substr)
		if index == -1 {
			break
		}

		count += 1
		word = word[index+len(substr):]
	}

	return count
}
