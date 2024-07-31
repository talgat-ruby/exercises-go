package main

import (
	"fmt"
	"strings"
)

func emojify(text string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if emoji, ok := emojiMap[word]; ok {
			words[i] = emoji
		}
	}

	return strings.Join(words, " ")
}

func main() {
	var c string
	fmt.Scan(&c)
	emojify(c)
}
