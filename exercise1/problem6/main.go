package main

import "strings"

func emojify(str string) string {
	emojis := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for key, emoji := range emojis {
		str = strings.ReplaceAll(str, key, emoji)
	}
	return str
}
