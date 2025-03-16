package main

import (
	"strings"
)

func emojify(text string) string {
	emojis := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for word, emoji := range emojis {
		text = strings.ReplaceAll(text, word, emoji)
	}
	return text
}
