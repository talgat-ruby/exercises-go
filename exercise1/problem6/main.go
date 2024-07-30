package main

import (
	"fmt"
	"strings"
)

func emojify(sentence string) string {
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	for word, emoji := range replacements {
		sentence = strings.Replace(sentence, word, emoji, -1)
	}

	return sentence
}

func main() {
	fmt.Println(emojify("Make me smile"))
	fmt.Println(emojify("Make me grin"))
	fmt.Println(emojify("Make me sad"))
	fmt.Println(emojify("Make me mad"))
}
