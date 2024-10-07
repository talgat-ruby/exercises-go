package main

import (
	"strings"
)

func potatoes(word string) int {
	return strings.Count(word, "potato")
}
