package main

import (
	"fmt"
	"strings"
)

// emojify function replaces specific words with their corresponding emojis
func emojify(s string) string {
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := strings.Split(s, " ")
	for i, word := range words {
		if emoji, exists := replacements[word]; exists {
			words[i] = emoji
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(emojify("Make me smile"))
	fmt.Println(emojify("Make me grin"))
	fmt.Println(emojify("Make me sad"))
	fmt.Println(emojify("Make me mad"))
}
