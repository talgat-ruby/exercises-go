package main

import (
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
		if emoji, ok := emojis[word]; ok {
			words[i] = emoji
		}
	}

	return strings.Join(words, " ")
}
