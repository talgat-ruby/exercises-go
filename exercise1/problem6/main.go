package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(emojify("Make me smile")) // "Make me 🙂"
	fmt.Println(emojify("Make me grin"))  // "Make me 😀"
	fmt.Println(emojify("Make me sad"))   // "Make me 😥"
}

func emojify(sentence string) string {
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for word, emoji := range replacements {
		sentence = strings.ReplaceAll(sentence, word, emoji)
	}
	return sentence
}
