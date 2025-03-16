package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)

	count := potatoes(text)
	fmt.Println("Number of potatoes in the string:", count)
}

// countPotatoes counts the number of occurrences of the word "potato" in a given string.
func potatoes(text string) int {
	return strings.Count(text, "potato")
}
