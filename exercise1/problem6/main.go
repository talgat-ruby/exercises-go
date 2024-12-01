package main

import (
	"fmt"
	"strings"
)

func emojify(sentence string) string {
	emojis := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	for word, emoji := range emojis {
		sentence = strings.ReplaceAll(sentence, word, emoji)
	}

	return sentence
}

func main() {
	fmt.Println(emojify("Make me smile")) // "Make me 🙂"
	fmt.Println(emojify("Make me grin"))  // "Make me 😀"
	fmt.Println(emojify("Make me sad"))   // "Make me 😥"
	fmt.Println(emojify("Make me mad"))   // "Make me 😠"
}
