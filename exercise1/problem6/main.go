package main

import "strings"

func emojify(str string) string {
	emojiMap := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}

	for word, emoji := range emojiMap {
		str = strings.Replace(str, word, emoji, -1)
	}

	return str
}
