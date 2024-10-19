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

	words := strings.Fields(sentence)

	for i, word := range words {
		cleanWord := strings.Trim(word, ",.!?;:")
		if emoji, exists := emojis[cleanWord]; exists {
			words[i] = strings.Replace(word, cleanWord, emoji, 1)
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(emojify("Make me smile"))
	fmt.Println(emojify("Make me grin"))
	fmt.Println(emojify("Make me sad"))
	fmt.Println(emojify("I am so mad!"))
}
