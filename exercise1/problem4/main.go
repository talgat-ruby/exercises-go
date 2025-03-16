package main

import (
	"fmt"
	"strings"
)

func main() {
	var crowd string
	fmt.Scan(&crowd)

	hiddenWord := detectWord(crowd)
	fmt.Println("Hidden word:", hiddenWord)
}

// findHiddenWord detects the hidden lowercase word in the crowd of uppercase letters.
func detectWord(crowd string) string {
	var hiddenWord strings.Builder

	for _, char := range crowd {
		if char >= 'a' && char <= 'z' {
			hiddenWord.WriteRune(char)
		}
	}

	return hiddenWord.String()
}
